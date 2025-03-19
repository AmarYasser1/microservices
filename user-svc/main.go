package main

import (
	"main/config"
	"main/initializer"
	"main/routes"
)



func main() {
	initializer.Init()

	r := routes.Routes()

	port := config.LoadConfig().ServerPort
    r.Run(":" + port)
}