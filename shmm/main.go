package main

import (
	"log"
	"net/http"
	"os"

	"github.com/freckie/shmfaas/shmm/endpoint"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Required .env file")
	}

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

	// Serve
	log.Println("Starting HTTP Server on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
