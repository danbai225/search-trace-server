package main

import (
	logs "github.com/danbai225/go-logs"
	"search-trace-server/config"
	"search-trace-server/db"
	"search-trace-server/http"
	"search-trace-server/server"
)

func main() {
	err := config.Load()
	if err != nil {
		logs.Err(err)
		return
	}
	err = db.InitDB()
	if err != nil {
		logs.Err(err)
		return
	}
	server.InitWordServer()
	http.Start()
}
