package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"learn-gin/app/pojo/req"
	"learn-gin/app/services"
	"learn-gin/config/log"
	"net/http"
)

type TestController struct {
}

func (t *TestController) HelloWorld(context *gin.Context) {
	log.Logger.Info("测试HelloWorld接口")
	context.String(http.StatusOK, "hello world")
}

func (t *TestController) TestParam(context *gin.Context) {
	name := context.Param("name")
	log.Logger.Info("测试TestParam接口")
	context.String(http.StatusOK, "check param %s", name)
}

func (t *TestController) TestDefaultParam(context *gin.Context) {
	name := context.DefaultQuery("name", "张三")
	gender := context.Query("gender")
	log.Logger.Info("测试TestDefaultParam接口")
	context.String(http.StatusOK, "他叫%s，性别:%s", name, gender)
}

func (t *TestController) TestPost(context *gin.Context) {
	name := context.PostForm("name")
	nick := context.DefaultPostForm("nick", "leo")
	log.Logger.Info("测试TestPost接口")
	context.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    http.StatusOK,
			"success": true,
		},
		"name": name,
		"nick": nick,
	})
}

func (t *TestController) TestPostBody(context *gin.Context) {
	var request req.TestPostRequest
	log.Logger.Info("测试TestPostBody接口")
	// 将前端穿过来的json数据绑定存储在这个实体类中，BindJSON()也能使用
	if err := context.ShouldBindJSON(&request); err != nil {
		log.Logger.Panic("参数异常")
	}

	if _, err := json.Marshal(request); err != nil {
		log.Logger.Panic("参数解析异常")
	}
	services.TestServ.PrintInfo(&request)

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": request,
	})
}
