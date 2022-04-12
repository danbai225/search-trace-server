package main

import (
	logs "github.com/danbai225/go-logs"
	"search-trace-server/config"
	"search-trace-server/http"
)

func main() {
	err := config.Load()
	if err != nil {
		logs.Err(err)
		return
	}
	http.Start()
}
