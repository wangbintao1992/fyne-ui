package config

import (
	"encoding/json"
	"io/ioutil"
)

type App struct {
	AppId  string   `json:"appId"`
	Title  string   `json:"title"`
	Child  []*Child `json:"child"`
	Weight float32  `json:"weight"`
	Height float32  `json:"height"`
}

type Child struct {
	Type      string    `json:"type"`
	Key       string    `json:"key"`
	isContext bool      `json:"isContext"`
	Order     int       `json:"order"`
	Child     []*Child  `json:"child"`
	Attr      *Attr     `json:"attr"`
	Function  *Function `json:"func"`
}

type Function struct {
	Type     string        `json:"type"`
	Script   string        `json:"script"`
	FuncName string        `json:"funcName"`
	Param    []interface{} `json:"param"`
}

type Attr struct {
	Text string `json:"text"`
	Row  int    `json:"row"`
}

func ParseConfig(path string) *App {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic("config file not found path:" + path)
	}

	app := &App{}
	err = json.Unmarshal(file, app)

	if err != nil {
		panic("config file init fail " + err.Error())
	}

	return app
}
