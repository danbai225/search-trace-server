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
		Debug  bool   `json:"debug"`
	} `json:"db"`
}

var C = Config{}
var f = ""

func Load() error {
	if f == "" {
		f = "conf.json"
	}
	env := os.Getenv("CONF_NAME")
	if env != "" {
		f = env
	}
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &C)
	if err != nil {
		return err
	}
	return nil
}
func LoadF(conf string) error {
	f = conf
	return Load()
}
