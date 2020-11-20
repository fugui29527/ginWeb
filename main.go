package main

import (
	"bytes"
	"fmt"
	"ginWeb/config"
	"ginWeb/log"
	"ginWeb/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

var logger *zap.Logger

func main() {
	app := gin.Default()
	app.Use(dolog())
	//初始化日志
	log.InitZapLog()
	logger = log.Logger
	//初始化路由
	initRouter(app)
	//初始化数据库连接
	_, dbErr := util.CreateOrmDb()
	if dbErr != nil {
		logger.Error("===数据库初始化失败:====", zap.Any("error:", dbErr))
	}
	//初始化redis
	util.CreateRedis()
	logger.Debug("========服务器启动成功==========")
	err := app.Run(":" + strconv.Itoa(config.AppConfig.AppInfo.Port))

	if err != nil {
		fmt.Println("......服务器启动失败..............")
	}

}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

//controller 的日志记录
func dolog() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		method := c.Request.Method
		url := c.Request.RequestURI
		logger.Info("=====客户端请求:", zap.String("方式method:", method), zap.String("路径url:", url))
		c.Next()
		logger.Info("=====服务端返回数据:", zap.String("方式method:", method),
			zap.String("路径url:", url), zap.String("数据data:", w.body.String()))
	}
}
