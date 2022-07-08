package services

import (
	"context"
	"learn-gin/app/pojo/rsp"
	"learn-gin/config/log"
	"learn-gin/config/redis"
	"time"
)

type RedisService interface {
	TestRedisInit() rsp.ResponseMsg
}

type RedisImpl struct {
}

//测试redis集群初始化使用
func (r RedisImpl) TestRedisInit() rsp.ResponseMsg {
	log.Logger.Info("测试redis集群初始化使用")
	_redis, _err := redis.GetRedisClusterClient()
	if _err != nil {
		log.Logger.Panic("获取redis连接异常", log.Any("error", _err.Error()))
	}
	_err = _redis.Set(context.TODO(), "haha", 100, 10*time.Second).Err()
	if _err != nil {
		log.Logger.Panic("设置缓存失败")
	}
	_v, _ := _redis.Get(context.TODO(), "haha").Result()

	log.Logger.Debug("haha缓存为:", log.String("v", _v))
	return *rsp.SuccessMsg("测试缓存")
}
