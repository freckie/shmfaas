package main

import (
	"net/http"
	"os"

	"github.com/freckie/shmfaas/shmm/endpoint"

	"github.com/julienschmidt/httprouter"
	klog "k8s.io/klog/v2"
)

type LoggingRouter struct {
	handler http.Handler
}

func (l *LoggingRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	klog.InfoSDepth(1, "Request arrived", "method", r.Method, "path", r.URL.Path)
	l.handler.ServeHTTP(w, r)
}

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
	logRouter := LoggingRouter{router}
	klog.InfoSDepth(0, "Starting HTTP Server", "port", port, "dbname", dbname)
	klog.ErrorDepth(0,
		http.ListenAndServe(":"+port, &logRouter),
		"Closing HTTP Server ...",
	)
}
