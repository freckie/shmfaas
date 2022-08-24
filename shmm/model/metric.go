package model

type ListMemResp struct {
	ShmSum     uint64        `json:"shm_sum"`
	ShmDisk    ListMemDevShm `json:"shm_disk"`
	ModelCount int           `json:"model_count"`
	Models     []ListMemItem `json:"models"`
}

type ListMemItem struct {
	ModelName string `json:"model_name"`
	TagName   string `json:"tag_name"`
	Shmname   string `json:"shmname"`
	Shmsize   uint64 `json:"shmsize"`
}

type ListMemDevShm struct {
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
	All  uint64 `json:"all"`
}
