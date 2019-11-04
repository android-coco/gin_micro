/*
 * @Author: yhlyl
 * @Date: 2019-11-03 11:06:12
 * @LastEditTime: 2019-11-04 21:21:47
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/model/user.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package model

import "gin_micro/db"

//CREATE TABLE `user` (
//`id` bigint(64) NOT NULL AUTO_INCREMENT,
//`mobile` varchar(20) DEFAULT NULL,
//`passwd` varchar(40) DEFAULT NULL,
//`avatar` varchar(150) DEFAULT NULL,
//`sex` varchar(2) DEFAULT NULL,
//`nickname` varchar(20) DEFAULT NULL,
//`salt` varchar(10) DEFAULT NULL,
//`online` int(10) DEFAULT NULL,
//`token` varchar(40) DEFAULT NULL,
//`memo` varchar(140) DEFAULT NULL,
//`createat` datetime DEFAULT NULL,
//`uid` varbinary(20) DEFAULT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
type User struct {
	Uid           int    `json:"uid"`
	Mobile        string `json:"mobile"`
	Token         string `json:"token"`
	LoginPassword string `json:"passwd"`
}

func (User) TableName() string {
	return "user"
}

func GetUserById(uid string) (User, error) {

	var user User
	err := db.GetDB().Where(" uid = ? ", uid).Find(&user).Error

	return user, err
}

func GetUserByAccount(account string) (User, error) {
	var user User
	err := db.GetDB().Where(" mobile = ? ", account).Find(&user).Error

	return user, err
}
