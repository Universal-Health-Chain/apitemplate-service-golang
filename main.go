package main

import (
	controllers "apitemplate-service-golang/controllers"
	"apitemplate-service-golang/utils/envUtils"
	"apitemplate-service-golang/utils/storageService"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Universal-Health-Chain/common-utils-golang/storageUtils"
)

// Note: the main function or test functions have to call envUtils.LoadEnv()

const (
	envHostHMAC = "INIT_HOST_EDV_HMAC"
	envHostDEK  = "INIT_HOST_EDV_DEK"
)

var (
	serviceName = os.Getenv("HOST_DB_NAME") // underscore instead of hyppen
	logPath     = "./logs/"                 // instead of the root "/"
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
	hostStorageParams := storageUtils.StorageParameters{
		StorageType:   storageUtils.DatabaseTypeMemOption,
		StorageURL:    "", // Empty for memory storage
		StoragePrefix: "prefix",
	}
	manager := &storageService.StorageServicesManager{}
	err := manager.CreateStorageService("host", hostStorageParams, 5)
	if err != nil {
		fmt.Println("Error creating storage service:", err)
		return
	}

	service, err := manager.GetStorageServiceByAlternateName("host")
	if err != nil {
		fmt.Println("Error getting storage service:", err)
		return
	}

	fmt.Println("Successfully retrieved storage service with alternate name:", service.GetAlternateName())

	// 4) Initializing the router
	envUtils.LoadEnv()
	InitLogger()

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
	storageService.InitHostDEK([]byte(os.Getenv(envHostDEK)))
	storageService.InitHostHMAC([]byte(os.Getenv(envHostHMAC)))
}
