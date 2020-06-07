package main

import (
	"github.com/tsubasahonda/go-app-boilerplate/config"
	"github.com/tsubasahonda/go-app-boilerplate/logs"
	"github.com/tsubasahonda/go-app-boilerplate/server"
)

func main() {
	c, err := config.ReadConfig()
	if err != nil {
		logs.Error("Invalid config: %s", err)
		panic(err)
	}

	sqlDB, err := connectToDB(c.DB.Host, c.DB.Port, c.DB.User, c.DB.Password, c.DB.DBName)
	if sqlDB != nil {
		defer sqlDB.Close()
	}
	if err != nil {
		logs.Error("DB connection failure: %s", err)
		panic(err)
	}
	if err = server.Init(c.Server.Port, sqlDB); err != nil {
		logs.Error("Failed starting server: %s", err)
		panic(err)
	}
}
