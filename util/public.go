package util

const (
	SuccessCode    = 0
	ErrorLackCode  = 1
	ErrorSqlCode   = 2
	ErrorRidesCode = 3

	//参数签名秘钥
	DesKey = "r5k1*8a$@8dc!dytkcs2dqz!"

	//redis key
	RedisKeyToken = "user:login:token:" //token 缓存

	//time
	FormatTime      = "15:04:05"            //时间格式
	FormatDate      = "2006-01-02"          //日期格式
	FormatDateTime  = "2006-01-02 15:04:05" //完整时间格式
	FormatDateTime2 = "2006-01-02 15:04"    //完整时间格式

	//微服务相关
	ConfigRoot      = "/gin_micro"
	RedisConfigRoot = ConfigRoot + "/redis/"
	RedisConfigPath = "config"
	MysqlConfigRoot = ConfigRoot + "/mysql/"
	MysqlConfigPath = "config"

	// 微服务名称
	GinMicroBb   = "gin_micro.service.db"
	GinMicroUser = "gin_micro.service.user"

	//队列名称
	UserQueue = "gin_micro_queue_user"
)
