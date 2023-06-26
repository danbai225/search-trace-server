package server

import (
	"encoding/json"
	"github.com/meilisearch/meilisearch-go"
	"search-trace-server/db"
	"search-trace-server/model"
)

type TraceServer interface {
	TraceCreate(trace *model.Trace) (err error)
	TraceSearchForKeyword(uName, key string, PageSize, PageNum int) (list []*model.Trace, PageTotal int64, total int64, err error)
}

type MySQLTraceServer struct{}
type MeiliSearchTraceServer struct{}

// TraceCreate 新建记录
func (MySQLTraceServer) TraceCreate(trace *model.Trace) (err error) {
	begin := db.GetDBW()
	defer func() {
		if err != nil {
			begin.Rollback()
		} else {
			begin.Commit()
		}
	}()
	err = begin.Create(trace).Error
	go WordParseNewWords(trace.Content)
	return err
}

// TraceSearchForKeyword 关键字搜索
func (MySQLTraceServer) TraceSearchForKeyword(uName, key string, PageSize, PageNum int) (list []*model.Trace, PageTotal int64, total int64, err error) {
	if PageSize == 0 {
		PageSize++
	}
	tx := db.GetDBR()
	defer tx.Commit()
	list = make([]*model.Trace, 0)
	err = tx.Limit(PageSize).Offset((PageNum-1)*PageSize).Model(&model.Trace{}).Where("username=? AND Match(title,content) Against (?)", uName, key).Scan(&list).Error
	err = tx.Model(&model.Trace{}).Where("username=? AND Match(title,content) Against (?)", uName, key).Count(&total).Error
	PageTotal = total / int64(PageSize)
	if total%int64(PageSize) != 0 {
		PageTotal++
	}
	return
}

// TraceCreate 新建记录
func (MeiliSearchTraceServer) TraceCreate(trace *model.Trace) (err error) {
	err = MySQLTraceServer{}.TraceCreate(trace)
	if err != nil {
		return err
	}
	index, err := db.GetMeiliSearchClient().GetIndex("trace")
	if err != nil {
		return err
	}
	_, err = index.AddDocuments(trace, "id")
	return err
}

// TraceSearchForKeyword 关键字搜索
func (MeiliSearchTraceServer) TraceSearchForKeyword(uName, key string, PageSize, PageNum int) (list []*model.Trace, PageTotal int64, total int64, err error) {
	list = make([]*model.Trace, 0)
	index, err := db.GetMeiliSearchClient().GetIndex("trace")
	if PageNum < 0 {
		PageNum = 1
	}
	of := int64((PageNum - 1) * PageSize)
	search, err := index.Search(key, &meilisearch.SearchRequest{
		Offset: of,
		Limit:  int64(PageSize),
		Filter: "username=" + uName,
	})
	if err != nil {
		return list, 0, 0, err
	}
	total = search.TotalHits
	marshalJSON, _ := json.Marshal(search.Hits)
	err = json.Unmarshal(marshalJSON, &list)
	return
}
