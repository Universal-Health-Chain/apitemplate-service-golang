package main

import (
	controllers "apitemplate-service-golang/controllers"
	"apitemplate-service-golang/utils/envUtils"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const logPath = "logs.log"

var Logger *log.Logger

func main() {
	envUtils.LoadEnv()
	InitLogger()

	// Initializing the Router for the API service
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	port := os.Getenv("HOST_LOCAL_PORT")
	if port == "" {
		port = "8024" //localhost
	}

	log.Println("Starting server on port " + port)
	log.Fatal(
		http.ListenAndServe(":"+port, controllers.CreateRouter()),
	)
}

func InitLogger() {

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open log file")
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

}
