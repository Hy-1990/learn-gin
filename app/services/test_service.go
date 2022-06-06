package services

import (
	"fmt"
	"learn-gin/app/pojo/req"
)

type TestService interface {
	PrintInfo(req *req.TestPostRequest)
}

type Test struct {
}

func (t Test) PrintInfo(req *req.TestPostRequest) {
	fmt.Printf("测试数据，name=%s,age=%d\n", req.Name, req.Age)
}
