package endpoint

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/freckie/shmfaas/shmm/entity"
	ihttp "github.com/freckie/shmfaas/shmm/internal/http"
	ishm "github.com/freckie/shmfaas/shmm/internal/posix_shm"
	"github.com/freckie/shmfaas/shmm/model"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/xid"
	"gorm.io/gorm"
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

// GET /shmodels/:name/:tag
func (e *Endpoint) GetModelTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := e.DB
	result := model.GetModelTagResp{}

	modelName := ps.ByName("name")
	if modelName == "" {
		ihttp.ResponseError(w, 404, "ModelName not found.")
		return
	}
	tagName := ps.ByName("tag")
	if tagName == "" {
		ihttp.ResponseError(w, 404, "TagName not found.")
		return
	}

	type queryResultType struct {
		Shmname   string
		Shmsize   int64
		CreatedAt time.Time
		UpdatedAt time.Time
		Metadata  []byte
	}
	var queryResult queryResultType

	dbResult := db.Model(&entity.SharedModel{}).
		Select("shmname, shmsize, created_at, updated_at, metadata").
		Where("name = ? AND tag = ?", modelName, tagName).
		First(&queryResult)
	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			ihttp.ResponseError(w, 404, "Model:Tag not found.")
		} else {
			ihttp.ResponseError(w, 500, dbResult.Error.Error())
		}
		return
	}

	result.ModelName = modelName
	result.TagName = tagName
	result.Shmname = queryResult.Shmname
	result.Shmsize = queryResult.Shmsize
	result.CreatedAt = queryResult.CreatedAt
	result.UpdatedAt = queryResult.UpdatedAt
	result.Metadata = queryResult.Metadata

	ihttp.ResponseOK(w, "Success", result)
}

// POST /shmodels/:name/:tag
func (e *Endpoint) PostModelTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := e.DB
	result := model.PostModelTagResp{}

	modelName := ps.ByName("name")
	if modelName == "" {
		ihttp.ResponseError(w, 404, "ModelName not found.")
		return
	}
	tagName := ps.ByName("tag")
	if tagName == "" {
		ihttp.ResponseError(w, 404, "TagName not found.")
		return
	}

	reqBody := model.PostModelTagReq{}
	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()
	err := decoder.Decode(&reqBody)
	if err != nil {
		ihttp.ResponseError(w, 400, "Invalid JSON body.")
		return
	}
	if reqBody.MemRequest <= 0 {
		ihttp.ResponseError(w, 400, "mem_request must be positive integer.")
		return
	}

	cntForValidation := 0
	dbResult := db.Model(&entity.SharedModel{}).
		Select("count(*) AS cnt").
		Where("name = ? AND tag = ?", modelName, tagName).
		First(&cntForValidation)
	if dbResult.Error != nil {
		if !errors.Is(dbResult.Error, gorm.ErrRecordNotFound) ||
			cntForValidation >= 1 {
			ihttp.ResponseError(w, 500, "Already exists.")
			return
		}
	}
	shmsize := reqBody.MemRequest
	shmname := xid.New().String()

	err = db.Transaction(func(tx *gorm.DB) error {
		_, err = ishm.Create(shmname, shmsize, 0666)
		if err != nil {
			return err
		}

		newModel := &entity.SharedModel{
			Name:     modelName,
			Tag:      tagName,
			Shmname:  shmname,
			Shmsize:  shmsize,
			Status:   0,
			Metadata: nil,
		}
		dbResult = tx.Model(&entity.SharedModel{}).
			Create(&newModel)
		if dbResult.Error != nil {
			return dbResult.Error
		}

		return nil
	})
	if err != nil {
		ihttp.ResponseError(w, 500, "Failed to create shm block : "+err.Error())
		return
	}

	result.Shmname = shmname
	result.Shmsize = shmsize
	ihttp.ResponseOK(w, "Success", result)
}

// DELETE /shmodels/:name/:tag
func (e *Endpoint) DeleteModelTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := e.DB

	modelName := ps.ByName("name")
	if modelName == "" {
		ihttp.ResponseError(w, 404, "ModelName not found.")
		return
	}
	tagName := ps.ByName("tag")
	if tagName == "" {
		ihttp.ResponseError(w, 404, "TagName not found.")
		return
	}

	var shmname string
	dbResult := db.Model(&entity.SharedModel{}).
		Select("shmname").
		Where("name = ? AND tag = ?", modelName, tagName).
		First(&shmname)
	if dbResult != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			ihttp.ResponseError(w, 404, "Model:Tag not found."+shmname)
			return
		}
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		err := ishm.Unlink(shmname)
		if err != nil {
			return err
		}

		dbResult := tx.Unscoped().
			Where("name = ? AND tag = ?", modelName, tagName).
			Delete(&entity.SharedModel{})
		if dbResult.Error != nil {
			return dbResult.Error
		}

		return nil
	})
	if err != nil {
		ihttp.ResponseError(w, 500, "Failed to delete shm block : "+err.Error())
		return
	}

	ihttp.ResponseOK(w, "Success", nil)
}
