package req

type StudentAddReq struct {
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Content string `json:"content"`
}

type StudentUpdateExecReq struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}
