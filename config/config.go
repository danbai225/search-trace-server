package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Db struct {
		Host   string `json:"host"`
		DbName string `json:"db_name"`
		User   string `json:"user"`
		Pass   string `json:"pass"`
		Port   int    `json:"port"`
	} `json:"db"`
}

var C = Config{}

func Load() error {
	f := "conf.json"
	env := os.ExpandEnv("CONF_NAME")
	if env != "" {
		f = env
	}
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	json.Unmarshal(file, &C)
	if err != nil {
		return err
	}
	return nil
}
