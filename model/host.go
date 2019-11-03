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
