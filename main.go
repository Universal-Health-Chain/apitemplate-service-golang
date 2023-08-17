package main

import (
	controllers "apitemplate-service-golang/controllers"
	"apitemplate-service-golang/utils/envUtils"
	"apitemplate-service-golang/utils/storageHostUtils"
	"apitemplate-service-golang/utils/storageProviderUtils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	storageProvider "github.com/trustbloc/edv/pkg/edvprovider/memedvprovider"
)

// Note: the main function or test functions have to call envUtils.LoadEnv()

const (
	envHostHMAC = "INIT_HOST_EDV_HMAC"
	envHostDEK  = "INIT_HOST_EDV_DEK"
)

var (
	serviceName = "apitemplate-service-golang"
	logPath     = "./logs/" // instead of the root "/"
	startTime   = fmt.Sprint(time.Now().Unix())
	logFile     = logPath + serviceName + "_" + startTime + ".txt"
	Logger      *log.Logger
)

// const envHostDBUrl = "HOST_DB_URL"

func main() {
	// 0) initializing the logger
	InitLogger()

	// 1) loading the environment variables
	envUtils.LoadEnv()

	// 2) Initializing host keys
	initHostKeys()

	// 3) Connecting to the host's storage
	storageProviderUtils.SetHostProviderMem(storageProvider.NewProvider())

	// 4) Initializing the router
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

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open log file")
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
}

func initHostKeys() {
	storageHostUtils.InitHostDEK([]byte(os.Getenv(envHostDEK)))
	storageHostUtils.InitHostHMAC([]byte(os.Getenv(envHostHMAC)))
}
