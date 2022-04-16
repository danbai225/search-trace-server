package server

import (
	"search-trace-server/db"
	"search-trace-server/model"
)

func AddTrace(trace *model.Trace) (err error) {
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
