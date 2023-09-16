package storageService

import (
	"apitemplate-service-golang/testData"
	"os"
	"testing"

	"github.com/Universal-Health-Chain/common-utils-golang/storageUtils"
)

func TestMongoDBStorage(t *testing.T) {
	InitHostDEK([]byte(os.Getenv("INIT_HOST_EDV_HMAC")))
	InitHostHMAC([]byte(os.Getenv("INIT_HOST_EDV_DEK")))

	opts := storageUtils.StorageParameters{
		StorageType:   storageUtils.DatabaseTypeMongoDBOption,
		StorageURL:    "mongodb://admin:password@localhost:27017",
		StoragePrefix: "prefix",
	}
	manager, err := NewStorageServicesManager("host", opts, 5)
	if err != nil {
		t.Fatal(err)
	}

	service, err := manager.GetStorageServiceByAlternateName("host")
	if err != nil {
		t.Fatal(err)
	}

	errMsg := service.HostNewClient([]byte(testData.TestResourceObjectStringifiedToInitStorage1), "host")
	if errMsg != "" {
		t.Fatal(errMsg)
	}
}

func TestCouchDBStorage(t *testing.T) {
	InitHostDEK([]byte(os.Getenv("INIT_HOST_EDV_HMAC")))
	InitHostHMAC([]byte(os.Getenv("INIT_HOST_EDV_DEK")))

	opts := storageUtils.StorageParameters{
		StorageType:   storageUtils.DatabaseTypeCouchDBOption,
		StorageURL:    "http://admin:password@localhost:5984",
		StoragePrefix: "prefix",
	}
	manager, err := NewStorageServicesManager("host", opts, 5)
	if err != nil {
		t.Fatal(err)
	}

	service, err := manager.GetStorageServiceByAlternateName("host")
	if err != nil {
		t.Fatal(err)
	}

	errMsg := service.HostNewClient([]byte(testData.TestResourceObjectStringifiedToInitStorage1), "host")
	if errMsg != "" {
		t.Fatal(errMsg)
	}
}
