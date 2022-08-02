package main

import (
	"net/http"
	"os"

	"github.com/freckie/shmfaas/shmm/endpoint"

	"github.com/julienschmidt/httprouter"
	klog "k8s.io/klog/v2"
)

func main() {
	// Environment variables
	port := os.Getenv("PORT")
	dbname := os.Getenv("SQLITE3")

	// DB Connection
	db, err := InitializeDB(dbname)
	if err != nil {
		panic("DB Connection failed")
	}
	ep := endpoint.Endpoint{DB: db}

	// Setting http router
	router := httprouter.New()
	router.GET("/shmodels", ep.ListSharedModel)
	router.GET("/shmodels/:name", ep.GetSharedModel)
	router.GET("/shmodels/:name/:tag", ep.GetModelTag)
	router.POST("/shmodels/:name/:tag", ep.PostModelTag)
	router.PUT("/shmodels/:name/:tag", ep.PutModelTag)
	router.DELETE("/shmodels/:name/:tag", ep.DeleteModelTag)
	router.GET("/shmodels/:name/:tag/accesses", ep.ListAccess)
	router.GET("/metrics/mem", ep.ListMem)
	router.GET("/metrics/health", ep.Health)

	// Serve
	klog.InfoSDepth(0, "Starting HTTP Server on port", port)
	klog.ErrorSDepth(0,
		http.ListenAndServe(":"+port, router),
		"Closing HTTP Server ...",
	)
}
