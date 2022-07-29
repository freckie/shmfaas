package model

import "time"

type ListAccessResp struct {
	Accesses    []ListAccessItem `json:"accesses"`
	AccessCount int              `json:"access_count"`
}

type ListAccessItem struct {
	ModelName  string    `json:"model_name"`
	TagName    string    `json:"tag_name"`
	AccessedAt time.Time `json:"accessed_at"`
}
