package main

import (
	"be/configs"
	"be/utils"
	"github.com/lithammer/shortuuid"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	log.Info(db)
	log.Info(shortuuid.New())
}
