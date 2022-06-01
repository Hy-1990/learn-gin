package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"learn-gin/config/log"
	"learn-gin/config/toml"
	"net/http"
)

type Result struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//反序列化为结构体对象
func parseJson(a string) Result {
	fmt.Printf("原始字符串: %s\n", a)
	var c Result
	if err := json.Unmarshal([]byte(a), &c); err != nil {
		fmt.Println("Error =", err)
		return c
	}
	return c
}

func main() {
	log.InitLogger(toml.GetConfig().Log.Path, toml.GetConfig().Log.Level)
	log.Logger.Info("hahahah")
	log.Logger.Info("config", log.Any("config", toml.GetConfig()))

	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})

	router.GET("/test/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "check param %s", name)
	})

	router.GET("/test1", func(context *gin.Context) {
		name := context.DefaultQuery("name", "张三")
		gender := context.Query("gender")
		context.String(http.StatusOK, "他叫%s，性别:%s", name, gender)
	})

	router.POST("/testPost", func(context *gin.Context) {
		name := context.PostForm("name")
		nick := context.DefaultPostForm("nick", "leo")
		context.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"code":    http.StatusOK,
				"success": true,
			},
			"name": name,
			"nick": nick,
		})
	})

	router.POST("/testPost2", func(context *gin.Context) {
		data, _ := ioutil.ReadAll(context.Request.Body)
		fmt.Println(string(data))
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": parseJson(string(data)),
		})
	})

	router.Run(":8080")
}
