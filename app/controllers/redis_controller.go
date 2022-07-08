package controllers

import (
	"github.com/gin-gonic/gin"
	"learn-gin/app/services"
	"learn-gin/config/log"
	"net/http"
)

type RedisController struct {
}

//测试redis初始化
func (s RedisController) TestRedisInit(context *gin.Context) {
	log.Logger.Info("测试redis初始化")
	_rsp := services.RedisServ.TestRedisInit()
	context.JSON(http.StatusOK, _rsp)
}
