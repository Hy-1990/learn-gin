package req

type StudentAddReq struct {
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Content string `json:"content"`
}
