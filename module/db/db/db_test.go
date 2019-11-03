package db

import (
	configWeb "gin_micro/config"
	"gin_micro/util"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/json"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/etcd"
	"testing"
)

func TestInitDB(t *testing.T) {
	eTCDSource := etcd.NewSource(
		// optionally specify etcd address; default to localhost:8500
		etcd.WithAddress(configWeb.GetService().EtcdUrl),
		// optionally specify prefix; defaults to /micro/config
		etcd.WithPrefix(util.MysqlConfigRoot),
		// optionally strip the provided prefix from the keys, defaults to false
		etcd.StripPrefix(true),
		source.WithEncoder(json.NewEncoder()),
	)
	// Create new config
	conf := config.NewConfig()

	// Load file source
	err := conf.Load(eTCDSource)
	if err != nil {
		t.Fatalf("微服务配置加载失败:%v", err)
	}
	var configDb configWeb.Db
	conf.Get(util.MysqlConfigPath)
	err = conf.Scan(&configDb)
	if err != nil {
		t.Fatalf("微服务启动失败:%v  %s", err, "读取配置文件")
	}
	_, err = InitDB(configDb)
	if err != nil {
		t.Fatalf("微服务启动失败:%v  %s", err, "数据库链接")
	}
}
