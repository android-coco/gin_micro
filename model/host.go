/*
 * @Author: yhlyl
 * @Date: 2019-11-03 12:21:08
 * @LastEditTime: 2019-11-04 21:21:39
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/model/host.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package model

import "github.com/jinzhu/gorm"

type QpHost struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

func (QpHost) TableName() string {
	return "host"
}

func GetHost(db *gorm.DB) ([]QpHost, error) {
	var hosts []QpHost

	err := db.Find(&hosts).Error

	if err != nil {
		return nil, err
	}

	return hosts, nil
}
