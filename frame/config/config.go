package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
	file, err := os.ReadFile(path)
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

func GetDefaultConfigPath() string {
	root, ok := getProjectRoot()

	if ok != nil {
		panic("config file not found path:" + root)
	}

	config := filepath.Join(root, "config.json")

	return config
}

func getProjectRoot() (string, error) {
	// 1. 获取当前执行代码的文件路径（开发/编译后都能正确获取）
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("获取当前文件路径失败")
	}

	// 2. 从当前文件路径向上遍历，找到包含go.mod的目录（项目根目录）
	dir := filepath.Dir(filePath)
	for {
		// 检查当前目录是否有go.mod
		goModPath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return dir, nil // 找到go.mod，返回该目录（项目根）
		}

		// 到达系统根目录仍未找到，终止遍历
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("未找到go.mod文件，非Go Module项目")
}
