package initializer

import (
	"main/config"
	"main/db"
)

var Cfg config.Config

func Init() {
	Cfg = config.LoadConfig()

	db.ConnectToDB()
	db.Migrate()
}