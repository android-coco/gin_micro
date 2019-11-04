/*
 * @Author: yhlyl
 * @Date: 2019-11-03 23:51:16
 * @LastEditTime: 2019-11-04 21:24:36
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/module/db/db/db_test.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package db

import (
	"gin_micro/module/config"
	"testing"
)

func TestInitDB(t *testing.T) {
	_, err := InitDB(config.GetDb())
	if err != nil {
		t.Fatalf("微服务启动失败:%v  %s", err, "数据库链接")
	}
}
