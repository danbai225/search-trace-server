package server

import (
	"search-trace-server/modle"
)

func AddTrace(trace *modle.Trace) error {
	begin := db.GetDB()
	err := begin.Create(trace).Error
	if err != nil {
		begin.Callback()
	} else {
		begin.Commit()
	}
	return err
}
