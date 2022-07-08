package router

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	// 测试路由
	TestRouter(r)
	// 学生路由
	StudentRouter(r)
	// 缓存路由
	RedisRouter(r)
}
