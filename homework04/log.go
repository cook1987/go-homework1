package main

import (
	"io"
	"net"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 全局日志实例
var logger *logrus.Logger

// 初始化日志系统
func initLogger() {
	logger = logrus.New()

	// 设置日志级别
	logger.SetLevel(logrus.InfoLevel)

	// 设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})

	// 同时输出到文件和终端
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Info("无法打开日志文件，使用标准输出")
	}

	// 添加钩子：同时输出到控制台
	logger.SetOutput(io.MultiWriter(os.Stdout, file))
}

// GinLogrusLogger Logrus 中间件
func GinLogrusLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		// 请求信息
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		// 日志字段
		fields := logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime.String(),
			"client_ip":    clientIP,
			"method":       reqMethod,
			"path":         reqUri,
			"user_agent":   c.Request.UserAgent(),
		}

		// 根据状态码记录不同级别的日志
		if statusCode >= 500 {
			logger.WithFields(fields).Error("服务器内部错误")
		} else if statusCode >= 400 {
			logger.WithFields(fields).Warn("客户端错误")
		} else {
			logger.WithFields(fields).Info("请求成功")
		}
	}
}

// GinLogrusRecovery 使用 Logrus 的 Recovery 中间件
func GinLogrusRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 检查连接是否已断开
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 获取请求信息
				clientIP := c.ClientIP()
				method := c.Request.Method
				path := c.Request.RequestURI

				// 记录 panic 信息
				logger.WithFields(logrus.Fields{
					"panic":     err,
					"client_ip": clientIP,
					"method":    method,
					"path":      path,
				}).Error("发生 panic")

				// 如果是连接断开，不尝试写入响应
				if brokenPipe {
					c.Error(err.(error))
					c.Abort()
				} else {
					c.AbortWithStatus(500)
				}
			}
		}()
		c.Next()
	}
}
