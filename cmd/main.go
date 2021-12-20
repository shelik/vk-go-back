package main

import (
	"log"

	"github.com/shelik/mtranslate/pkg/config"
	"github.com/shelik/mtranslate/server"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}
	app := server.NewApp()
	app.Run(viper.GetString("port"))
}
