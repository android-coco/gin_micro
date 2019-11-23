/*
 * @Author: yhlyl
 * @Date: 2019-11-03 10:56:49
 * @LastEditTime: 2019-11-06 15:33:38
 * @LastEditors: yhlyl
 * @Description: In User Settings Edit
 * @FilePath: /gin_micro/cmd/main.go
 */
package main

import (
	redis "gin_micro/cache"
	"gin_micro/config"
	"gin_micro/db"
	"gin_micro/httpserver"
	"gin_micro/tcp"
	"log"
)

func main() {
	config.InitConfig("/../config/config.yml")
	err := redis.InitRedis(config.GetRedis())
	if err != nil {
		log.Fatalf("redis init err %v", err)
		return
	}
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
