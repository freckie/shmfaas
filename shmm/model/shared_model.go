package model

type ListSharedModelResp struct {
	SharedModels     []ListSharedModelItem `json:"shared_models"`
	SharedModelCount int                   `json:"shared_model_count"`
}

type ListSharedModelItem struct {
	ModelName string `json:"model_name"`
	TagCount  int    `json:"tag_count"`
}

type GetSharedModelResp struct {
	ModelName string   `json:"model_name"`
	Tags      []string `json:"tags"`
	TagCount  int      `json:"tag_count"`
}

type GetModelTagResp struct {
	ModelName string `json:"model_name"`
	TagName   string `json:"tag_name"`
	Shmname   string `json:"shmname"`
	Shmsize   int64  `json:"shmsize"`
	Metadata  string `json:"metadata"`
}

type PostModelTagReq struct {
	MemRequest int64 `json:"mem_request"`
}

type PostModelTagResp struct {
	Shmname string `json:"shmname"`
	Shmsize int64  `json:"shmsize"`
}

type PutModelTagReq struct {
	Metadata string `json:"metadata"`
}

type PutModelTagResp struct {
	ModelName string `json:"model_name"`
	TagName   string `json:"tag_name"`
	Shmname   string `json:"shmname"`
	Shmsize   int64  `json:"shmsize"`
	Metadata  string `json:"metadata"`
}
