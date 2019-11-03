package main

import (
	"gin_micro/cache"
	"gin_micro/config"
	"gin_micro/db"
	"gin_micro/httpserver"
	"gin_micro/tcp"
	"log"
)

func main() {
	config.InitConfig("/../config/config.yml")
	redis.InitRedis(config.GetRedis())
	initDB, err := db.InitDB(config.GetDb())
	if err != nil {
		log.Fatalf("db init err %v", err)
		return
	}
	defer func() {
		err := initDB.Close()
		if err != nil {
			log.Fatalf("db close err %v", err)
		}
	}()
	router := httpserver.SetupRouter()
	go tcp.Run(config.GetService().TCPPort)
	err = router.Run(config.GetService().Port)
	if err != nil {
		log.Fatalf("http server run err %v", err)
	}

}
