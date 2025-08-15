package inits

func init() {
	InitViper()   //初始化动态配置
	InitNaCos()   //初始化线上配置
	InitMysql()   //初始化数据库
	InitRedis()   //初始化redis缓存
	InitEs()      //初始化es客户端
	InitMongoDB() //初始化mongodb数据库
}
