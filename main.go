package main

import (
	"be-park-ease/config"
	"be-park-ease/logger"
	"fmt"
)

func init() {
	config.Init()
	conf := config.Get()
	logger.Init(conf)
	logger.UpdateLogLevel(conf.Logger.Level)
}

func main() {
	fmt.Println("Hello World!")
}
