package main

import "learn-gin-mvc/application/controller"

func main() {
	router := controller.GetRouter()
	router.Run(":8001")
}
