package main

import "mvc-model-sample/application/controller"

func main() {
	router := controller.GetRouter()
	router.Run(":8001")
}
