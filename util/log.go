package util

import (
	"os"
	"path/filepath"

	"github.com/cihub/seelog"
)

const ConfigDefaultLogConfigFile = "/../config/log.xml"

var Logger seelog.LoggerInterface

//func InitLog(logConfigFile string) error {
//	defer seelog.Flush()
//	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
//	if logConfigFile == "" {
//		logConfigFile = ConfigDefaultLogConfigFile
//	}
//	//初始化全局变量Logger为seelog的禁用状态，主要为了防止Logger被多次初始化
//	var err error
//	_ = seelog.ReplaceLogger(Logger)
//	Logger, err = seelog.LoggerFromConfigAsFile(path + logConfigFile)
//	return err
//}

func init() {
	defer seelog.Flush()
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//初始化全局变量Logger为seelog的禁用状态，主要为了防止Logger被多次初始化
	_ = seelog.ReplaceLogger(Logger)
	Logger, _ = seelog.LoggerFromConfigAsFile(path + ConfigDefaultLogConfigFile)
}
