package main

import (
	"main/initializers"
	"main/routes"
)

func main() {
	initializer.Init()

	r := routes.Routes()

	port := initializer.Cfg.ServerPort
    r.Run(":" + port)
}


/*
    "id"
    "userId"
    "product_name"
    "quantity"
*/