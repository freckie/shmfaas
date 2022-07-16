package endpoint

import (
	"net/http"

	"github.com/freckie/shmfaas/shmm/entity"
	ihttp "github.com/freckie/shmfaas/shmm/internal/http"
	"github.com/freckie/shmfaas/shmm/model"

	"github.com/julienschmidt/httprouter"
)

// GET /metric/health
func (e *Endpoint) Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ihttp.ResponseOK(w, "Health", nil)
}

// GET /metric/mem
func (e *Endpoint) ListMem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := e.DB
	result := model.ListMemResp{}

	var queryResult []model.ListMemItem
	dbResult := db.Model(&entity.SharedModel{}).
		Select("name AS model_name, tag AS tag_name, shmname, shmsize").
		Find(&queryResult)
	if dbResult.Error != nil {
		ihttp.ResponseError(w, 500, dbResult.Error.Error())
		return
	}

	result.MemSum = 0
	result.Models = queryResult
	for _, iter := range queryResult {
		result.MemSum += iter.Shmsize
	}
	result.ModelCount = len(queryResult)

	ihttp.ResponseOK(w, "Success", result)
}
