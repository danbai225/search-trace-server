package server

import (
	"search-trace-server/db"
	"search-trace-server/model"
)

// TraceCreate 新建记录
func TraceCreate(trace *model.Trace) (err error) {
	begin := db.GetDB()
	defer func() {
		if err != nil {
			begin.Callback()
		} else {
			begin.Commit()
		}
	}()
	err = begin.Create(trace).Error
	return err
}
