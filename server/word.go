package server

import (
	"context"
	"fmt"
	logs "github.com/danbai225/go-logs"
	"github.com/go-ego/gse"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/yanyiwu/gojieba"
	"search-trace-server/config"
	"search-trace-server/db"
	"search-trace-server/model"
	"strings"
	"time"
)

const (
	wordCachePrefix = "word_"
	wordCacheTime   = time.Hour * 24
)

var (
	seg      gse.Segmenter
	ctx      = gctx.New()
	setWords = gset.New(true)
)

func InitWordServer() {
	go func() {
		if !config.C.LiteMode {
			err := seg.LoadDictEmbed()
			if err != nil {
				logs.Err(err)
				return
			}
			err = seg.LoadStopEmbed()
			if err != nil {
				logs.Err(err)
				return
			}
		}
		gtimer.Add(ctx, time.Minute, func(ctx context.Context) {
			slice := setWords.Slice()
			setWords.Clear()
			arr := make([]string, 0)
			for _, w := range slice {
				arr = append(arr, w.(string))
			}
			WordCreateList(arr)
		})
	}()
}

// WordCreate 创建分词
func WordCreate(word string) (err error) {
	tx := db.GetDBW()
	defer func() {
		if err != nil {
			tx.Callback()
			if strings.Contains(err.Error(), "Duplicate") {
				db.GetCache().Add(fmt.Sprint(wordCachePrefix, word), wordCacheTime, true)
			}
		} else {
			tx.Commit()
		}
	}()
	err = tx.Create(&model.Word{Word: word}).Error
	return
}
func WordCreateList(words []string) {
	tx := db.GetDBW()
	defer func() {
		tx.Commit()
	}()
	for _, word := range words {
		err := tx.Create(&model.Word{Word: word}).Error
		if err != nil {
			tx.Callback()
			if strings.Contains(err.Error(), "Duplicate") {
				db.GetCache().Add(fmt.Sprint(wordCachePrefix, word), wordCacheTime, true)
			}
		}
	}
	return
}

// WordParseNewWords 解析新词
func WordParseNewWords(text string) {
	cache := db.GetCache()
	var search []string
	if config.C.LiteMode {
		x := gojieba.NewJieba()
		defer x.Free()
		search = x.CutForSearch(text, true)
	} else {
		search = seg.CutSearch(text, true)
	}
	words := make([]string, 0)
	for _, w := range search {
		if len(w) > 4 {
			if !cache.Exists(fmt.Sprint(wordCachePrefix, w)) {
				//_ = WordCreate(w)
				words = append(words, w)
				cache.Add(fmt.Sprint(wordCachePrefix, w), wordCacheTime, true)
			}
		}
	}
	for _, word := range words {
		setWords.Add(word)
	}
}

// WordAssociate 关键字联想
func WordAssociate(word string) (list []string, err error) {
	tx := db.GetDBR()
	defer tx.Commit()
	list = make([]string, 0)
	err = tx.Limit(10).Offset(0).Model(&model.Word{}).Where("word like ? AND word !=?", fmt.Sprint(word, "%"), word).Scan(&list).Error
	return
}
