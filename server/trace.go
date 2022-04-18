package server

import (
	"search-trace-server/db"
	"search-trace-server/model"
)

// TraceCreate 新建记录
func TraceCreate(trace *model.Trace) (err error) {
	begin := db.GetDBW()
	defer func() {
		if err != nil {
			begin.Callback()
		} else {
			begin.Commit()
		}
	}()
	err = begin.Create(trace).Error
	go WordParseNewWords(trace.Content)
	return err
}

// TraceSearchForKeyword 关键字搜索
func TraceSearchForKeyword(uName, key string, PageSize, PageNum int) (list []*model.Trace, PageTotal int64, err error) {
	if PageSize == 0 {
		PageSize++
	}
	tx := db.GetDBR()
	list = make([]*model.Trace, 0)
	err = tx.Limit(PageSize).Offset((PageNum-1)*PageSize).Model(&model.Trace{}).Where("username=? AND Match(title,content) Against (?)", uName, key).Scan(&list).Error
	total := int64(0)
	err = tx.Model(&model.Trace{}).Where("username=? AND Match(title,content) Against (?)", uName, key).Count(&total).Error
	PageTotal = total / int64(PageSize)
	if total%int64(PageSize) != 0 {
		PageTotal++
	}
	return
}
