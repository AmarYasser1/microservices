package main

import (
	"main/initializer"
	"main/routes"
)

func main() {
	initializer.Init()

	r := routes.Routes()

	port := initializer.Cfg.ServerPort
    r.Run(":" + port)
}


/*
   "id" : 
   "name" :
   "email" :
   "password" :
*/