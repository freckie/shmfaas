package endpoint

import (
	"net/http"
	"syscall"

	"github.com/freckie/shmfaas/shmm/entity"
	ihttp "github.com/freckie/shmfaas/shmm/internal/http"
	"github.com/freckie/shmfaas/shmm/model"

	"github.com/julienschmidt/httprouter"
)

// GET /metric/health
func (e *Endpoint) Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ihttp.ResponseOK(w, r, "Health", nil)
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
		ihttp.ResponseError(w, r, 500, dbResult.Error.Error())
		return
	}

	result.ShmSum = 0
	result.Models = queryResult
	for _, iter := range queryResult {
		result.ShmSum += iter.Shmsize
	}
	result.ModelCount = len(queryResult)

	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/dev/shm", &fs)
	if err != nil {
		ihttp.ResponseError(w, r, 500, "Cannot get the status of /dev/shm : "+err.Error())
		return
	}
	result.ShmDisk.All = fs.Blocks * uint64(fs.Bsize)
	result.ShmDisk.Free = fs.Bfree * uint64(fs.Bsize)
	result.ShmDisk.Used = result.ShmDisk.All - result.ShmDisk.Free

	ihttp.ResponseOK(w, r, "Success", result)
}
