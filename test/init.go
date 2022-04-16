package test

import (
	logs "github.com/danbai225/go-logs"
	"search-trace-server/config"
	"search-trace-server/db"
)

func init() {
	err := config.LoadF("../conf-dev.json")
	if err != nil {
		logs.Err(err)
		return
	}
	err = db.InitDB()
	if err != nil {
		logs.Err(err)
		return
	}
}
