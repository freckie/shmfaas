package model

type ListMemResp struct {
	MemSum     int64         `json:"mem_sum"`
	ModelCount int           `json:"model_count"`
	Models     []ListMemItem `json:"models"`
}

type ListMemItem struct {
	ModelName string `json:"model_name"`
	TagName   string `json:"tag_name"`
	Shmname   string `json:"shmname"`
	Shmsize   int64  `json:"shmsize"`
}
