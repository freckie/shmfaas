package endpoint

import (
	"errors"
	"net/http"
	"time"

	"github.com/freckie/shmfaas/shmm/entity"
	ihttp "github.com/freckie/shmfaas/shmm/internal/http"
	"github.com/freckie/shmfaas/shmm/model"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

// GET /shmodels/:name/:tag/accesses
func (e *Endpoint) ListAccess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := e.DB
	result := model.ListAccessResp{}

	modelName := ps.ByName("name")
	if modelName == "" {
		ihttp.ResponseError(w, r, 404, "ModelName not found.")
		return
	}
	tagName := ps.ByName("tag")
	if tagName == "" {
		ihttp.ResponseError(w, r, 404, "TagName not found.")
		return
	}

	type queryResultType struct {
		ModelName  string
		ModelTag   string
		AccessedAt time.Time
	}
	var queryResult []queryResultType

	dbResult := db.Model(&entity.Access{}).
		Select("model_name, model_tag, accessed_at").
		Where("model_name = ? AND model_tag = ?", modelName, tagName).
		Order("accessed_at DESC").
		Find(&queryResult)
	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			ihttp.ResponseError(w, r, 404, "Model:Tag not found.")
		} else {
			ihttp.ResponseError(w, r, 500, dbResult.Error.Error())
		}
		return
	}

	result.Accesses = make([]model.ListAccessItem, len(queryResult))
	for idx, item := range queryResult {
		result.Accesses[idx].ModelName = item.ModelName
		result.Accesses[idx].TagName = item.ModelTag
		result.Accesses[idx].AccessedAt = item.AccessedAt
	}
	result.AccessCount = len(result.Accesses)

	ihttp.ResponseOK(w, r, "Success", result)
}
