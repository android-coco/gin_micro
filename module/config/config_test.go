package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	public := GetPublic()
	db := GetDb()
	redis := GetRedis()
	fmt.Println(public)
	fmt.Println(db)
	fmt.Println(redis)
}