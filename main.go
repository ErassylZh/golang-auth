package main

import (
	"auth/app"
	"auth/config"
)

func main() {

	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
