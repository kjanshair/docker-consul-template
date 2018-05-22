package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Version string `json:"version"`
	AppName string `json:"appname"`
}

func main() {
	file, err := os.Open("config/config.json")
	if err != nil {
		fmt.Print(err)
	}
	robots, err := ioutil.ReadAll(file)

	var config Config
	json.Unmarshal([]byte(robots), &config)

	fmt.Println(config.Version)
	fmt.Println(config.AppName)

	defer file.Close()
}

