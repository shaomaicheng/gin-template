package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/donnie4w/go-logger"
)

var DB = make(map[string]string)

func initLogger() {
	if gin.IsDebugging() {
		//指定是否控制台打印，默认为true
		logger.SetConsole(false)
		//指定日志文件备份方式为文件大小的方式
		//第一个参数为日志文件存放目录
		//第二个参数为日志文件命名
		//第三个参数为备份文件最大数量
		//第四个参数为备份文件大小
		//第五个参数为文件大小的单位 KB，MB，GB TB
		//logger.SetRollingFile("d:/logtest", "test.log", 10, 5, logger.KB)

		//指定日志文件备份方式为日期的方式
		//第一个参数为日志文件存放目录
		//第二个参数为日志文件命名
		logger.SetRollingDaily("./loggers", "gin.log")

		//指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
		//一般习惯是测试阶段为debug，生成环境为info以上
		logger.SetLevel(logger.DEBUG)
	} else {
		logger.SetConsole(false)

		//指定日志文件备份方式为日期的方式
		//第一个参数为日志文件存放目录
		//第二个参数为日志文件命名
		logger.SetRollingDaily("./loggers", "gin.log")

		//指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
		//一般习惯是测试阶段为debug，生成环境为info以上
		logger.SetLevel(logger.INFO)
	}
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		logger.Debug(user)
		value, ok := DB[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}


		if c.Bind(&json) == nil {
			logger.Debug("just for test")
			DB[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	initLogger()
	logger.Debug("init log finished!")

	r := setupRouter()
	logger.Debug("init router finished!")
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
