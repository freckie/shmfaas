package shmscoring

type metricsResp struct {
	ShmSum     int64               `json:"shm_sum"`
	ShmDisk    metricsRespShmDisk  `json:"shm_disk"`
	ModelCount int                 `json:"model_count"`
	Models     []metricsRespModels `json:"models"`
}

type metricsRespShmDisk struct {
	Used int64 `json:"used"`
	Free int64 `json:"free"`
	All  int64 `json:"all"`
}

type metricsRespModels struct {
	ModelName string `json:"model_name"`
	TagName   string `json:"tag_name"`
}

type NodeMetric struct {
	ShmSum      int64
	ShmDiskUsed int64
	ShmDiskFree int64
	ShmDiskAll  int64
	ModelCount  int
	Models      []string
}
