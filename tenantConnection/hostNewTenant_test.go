package tenantConnection

import (
	"connections-service/testData"
	"encoding/base64"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/contentUtils"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/edvUtils"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/edvUtilsTest"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/storageHostUtils"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/tenantUtils"
	ariesMongo "github.com/hyperledger/aries-framework-go-ext/component/storage/mongodb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hostNewTenant(t *testing.T) {
	// Mocked function to generate a fixed profile code for a concrete user (shall be different for every profile)
	oriNewUUIDv4Base58 := contentUtils.NewUUIDv4Base58
	contentUtils.NewUUIDv4Base58 = func() string {
		return TestProfileActivationCodeFromProviderForDirectorICT
	}

	storageHostUtils.InitHostDEK([]byte("anything"))
	storageHostUtils.InitHostHMAC([]byte("anything"))

	// 1) starting the DB server for the Host
	dbServerHostConnectionStr := testData.TestConnectionsServiceMongoConnStrWithPortHost
	dbServerHostMongoProvider, err := ariesMongo.NewProvider(dbServerHostConnectionStr)
	assert.NoError(t, err)

	dbHostDockerPool, dbHostDockerPod := edvUtilsTest.DockerStartMongoDBContainer(t, dbServerHostMongoProvider, testData.TestPortMongoServerForConnectionsServiceHost)
	// defer function is defined here, but it is executed when the test has finished
	defer func() {
		assert.NoError(t, dbHostDockerPool.Purge(dbHostDockerPod), "failed to purge host's MongoDB resource in Docker")
	}()

	// creating the EDV provider by using the MongoDB Host provider
	edvHostProvider, err := edvUtils.NewProvider(dbServerHostMongoProvider, "configurations", "documents", 100)
	assert.NoError(t, err)
	storageHostUtils.SetProviderHostEDV(edvHostProvider)
	edvHostProviderRecovered := storageHostUtils.GetHostProviderEDV()
	assert.NotNil(t, edvHostProviderRecovered)

	// 2) starting the DB server for the first tenant
	dbConnectionStrServerTenant1 := testData.TestConnectionsServiceMongoConnStrWithPortTenant1
	dbMongoProviderServerTenant1, err := ariesMongo.NewProvider(dbConnectionStrServerTenant1)
	assert.NoError(t, err)

	dbDockerPoolTenant1, dbDockerPodTenant1 := edvUtilsTest.DockerStartMongoDBContainer(t, dbMongoProviderServerTenant1, testData.TestPortMongoServerForConnectionsServiceTenant1)
	assert.NoError(t, err)
	// defer function is defined here, but it is executed when the test has finished
	defer func() {
		assert.NoError(t, dbDockerPoolTenant1.Purge(dbDockerPodTenant1), "failed to purge second tenant's MongoDB resource in Docker")
	}()

	initHostEnvBase64 := base64.RawURLEncoding.EncodeToString([]byte(testData.TestInitConnectionsServiceWithTenantStringified))
	initHostEnvBytes, _ := base64.RawURLEncoding.DecodeString(initHostEnvBase64)

	// HostNewTenant stores the TenantBackend encrypted in the Host's EDV
	resObjArray, errObjArray := HostNewTenantInConnectionService(initHostEnvBytes)
	assert.NotNil(t, resObjArray)
	assert.Empty(t, errObjArray)

	// GetTenantBackendInHostEDV gets and decrypts the TenantBackend stored in the Host's EDV
	tenantBackend, errMsg := storageHostUtils.GetTenantBackendInHostEDV(edvHostProvider, tenantUtils.TestAlternateNameValueTenant1)
	assert.NotNil(t, tenantBackend)
	assert.Empty(t, errMsg)

	contentUtils.NewUUIDv4Base58 = oriNewUUIDv4Base58
}
