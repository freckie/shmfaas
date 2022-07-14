package model

import "time"

type ListSharedModelResp struct {
	SharedModels     []ListSharedModelItem `json:"shared_models"`
	SharedModelCount int                   `json:"shared_model_count"`
}

type ListSharedModelItem struct {
	ModelName string `json:"model_name"`
	TagCount  int    `json:"tag_count"`
}

type GetSharedModelResp struct {
	ModelName string               `json:"model_name"`
	Tags      []GetSharedModelItem `json:"tags"`
	TagCount  int                  `json:"tag_count"`
}

type GetSharedModelItem struct {
	TagName   string    `json:"tag_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetModelTagResp struct {
	ModelName string    `json:"model_name"`
	TagName   string    `json:"tag_name"`
	Shmname   string    `json:"shmname"`
	Shmsize   int64     `json:"shmsize"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Metadata  []byte    `json:"metadata"`
}

type PostModelTagReq struct {
	MemRequest int64 `json:"mem_request"`
}

type PostModelTagResp struct {
	Shmname string `json:"shmname"`
	Shmsize int64  `json:"shmsize"`
}

type PutModelTagReq struct {
	Shmname  string `json:"shmname"`
	Shmsize  int64  `json:"shmsize"`
	Metadata []byte `json:"metadata"`
}

type PutModelTagResp struct {
	ModelName string    `json:"model_name"`
	TagName   string    `json:"tag_name"`
	Shmname   string    `json:"shmname"`
	Shmsize   int64     `json:"shmsize"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Metadata  []byte    `json:"metadata"`
}
