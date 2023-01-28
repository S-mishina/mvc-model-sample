package main

import (
	"mvc-model-sample/application/controller"
	"mvc-model-sample/application/config"
)

func main() {
	config.LoadEnv()
	router := controller.GetRouter()
	router.Run(":8001")
}
