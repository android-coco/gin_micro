/*
 * @Author: yhlyl
 * @Date: 2019-11-03 10:57:45
 * @LastEditTime: 2019-11-04 15:56:54
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /gin_micro/db/db.go
 */
package db

import (
	"fmt"
	"gin_micro/config"
	jasonlog "gin_micro/util/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type dbLogger struct {
}

func (logger dbLogger) Print(values ...interface{}) {
	jasonlog.Info(values)
}

// Connect 初始化 DB连接
func InitDB(configDb config.Db) (*gorm.DB, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configDb.User, configDb.PassWd, configDb.Host, configDb.Db)
	var err error
	db, err = gorm.Open(configDb.Dialect, args)
	if err != nil {
		fmt.Printf("Can't connect to database, dialect: %s, args: %s err: %v\n", configDb.Dialect, args, err)
		return nil, err
	}
	db.LogMode(configDb.EnableLog)
	db.SetLogger(dbLogger{})
	//用于设置最大打开的连接数，默认值为0表示不限制。
	db.DB().SetMaxOpenConns(configDb.MaxOpenConnections)
	//用于设置闲置的连接数。
	db.DB().SetMaxIdleConns(configDb.MaxIdleConnections)

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
