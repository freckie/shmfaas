module github.com/freckie/shmfaas/shmm

go 1.18

require (
	github.com/joho/godotenv v1.4.0
	github.com/julienschmidt/httprouter v1.3.0
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.8
)

require (
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.12 // indirect
	github.com/rs/xid v1.4.0 // indirect
	k8s.io/klog/v2 v2.70.1 // indirect
)

replace (
	github.com/freckie/shmfaas/shmm/endpoint => ./endpoint
	github.com/freckie/shmfaas/shmm/entity => ./entity
	github.com/freckie/shmfaas/shmm/internal/http => ./internal/http
	github.com/freckie/shmfaas/shmm/internal/posix_shm => ./internal/posix_shm
	github.com/freckie/shmfaas/shmm/model => ./model
)
