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
