package main

import (
	"log"

	"github.com/shelik/vk-go-back/pkg/config"
	"github.com/shelik/vk-go-back/server"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}
	app := server.NewApp()
	app.Run(viper.GetString("port"))
}
