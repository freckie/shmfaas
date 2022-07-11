package endpoint

import (
	"net/http"
	"time"

	"github.com/freckie/shmfaas/shmm/entity"
	ihttp "github.com/freckie/shmfaas/shmm/internal/http"
	"github.com/freckie/shmfaas/shmm/model"

	"github.com/julienschmidt/httprouter"
)

// GET /shmodels
func (e *Endpoint) ListSharedModel(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := e.DB
	result := model.ListSharedModelResp{}

	type queryResultType struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}
	var queryResult []queryResultType

	dbResult := db.Model(&entity.SharedModel{}).
		Select("name, count(tag) as count").
		Group("name").
		Find(&queryResult)
	if dbResult.Error != nil {
		ihttp.ResponseError(w, 500, dbResult.Error.Error())
		return
	}

	result.SharedModelCount = len(queryResult)
	result.SharedModels = make([]model.ListSharedModelItem, len(queryResult))
	for idx, iter := range queryResult {
		result.SharedModels[idx].ModelName = iter.Name
		result.SharedModels[idx].TagCount = iter.Count
	}

	ihttp.ResponseOK(w, "Success", result)
}

// GET /shmodels/:name
func (e *Endpoint) GetSharedModel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := e.DB
	result := model.GetSharedModelResp{}

	modelName := ps.ByName("name")
	if modelName == "" {
		ihttp.ResponseError(w, 404, "ModelName not found.")
		return
	}

	type queryResultType struct {
		Name      string    `json:"name"`
		Tag       string    `json:"tag"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	var queryResult []queryResultType

	dbResult := db.Model(&entity.SharedModel{}).
		Select("name, tag, created_at, updated_at").
		Where("name = ?", modelName).
		Find(&queryResult)
	if dbResult.Error != nil {
		ihttp.ResponseError(w, 500, dbResult.Error.Error())
		return
	}

	if len(queryResult) <= 0 {
		ihttp.ResponseError(w, 404, "ModelName not found.")
		return
	}

	result.ModelName = modelName
	result.TagCount = len(queryResult)
	result.Tags = make([]model.GetSharedModelItem, len(queryResult))
	for idx, iter := range queryResult {
		result.Tags[idx].TagName = iter.Tag
		result.Tags[idx].CreatedAt = iter.CreatedAt
		result.Tags[idx].UpdatedAt = iter.UpdatedAt
	}

	ihttp.ResponseOK(w, "Success", result)
}
