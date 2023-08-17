package tenantConnection

import (
	"apitemplate-service-golang/edvConnection"
	"encoding/json"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/didcommUtils"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/storageHostUtils"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/tenantUtils"
	"net/http"
	"strconv"
	"strings"
)

/*
func ReadAndDecodeTenantFromEnv() []byte {
	envTenantBase64 := os.Getenv("INIT_TENANT")
	envDecodedTenantBytes := base64Utils.DecodeBase64UrlToBytes(envTenantBase64)
	//fmt.Println(envDecodedTenantBytes)
	return envDecodedTenantBytes
}*/

// HostNewTenant gets employees and departments, creates administrative departments if required and initializes a DEK for the tenant,
// so EDV data can be stored in the tenant's DB service, encrypted for the tenant's public storage encryption key (the DEK).
// Admin users in the tenant's DB service can only see encrypted EDVs,
// and the tenant's DEK remains encrypted in the Host, so admin users in the Host's DB cannot see the DEK of the tenant.
// Devops users in the Host can only see a seed for the Host DEK,
// but they are not able to derive the DEK of the host to decrypt the tenant's DB connection (URL, user and password),
// nor the seed of the tenant's DEK to decrypt EDV data stored in the tenant's DB service.
var HostNewTenantInConnectionService = func(resourceObjectBytes []byte) (resObjArray []didcommUtils.ResourceObject, errObjArray []didcommUtils.ErrorObject) {

	var tenantResourceObject didcommUtils.ResourceObject
	err := json.Unmarshal(resourceObjectBytes, &tenantResourceObject)
	if err != nil {
		errMsg := "error unmarshalling the resource object of the tenant"
		errCode := strconv.Itoa(http.StatusBadRequest)
		errorObjs := []didcommUtils.ErrorObject{{
			Code:  &errCode,
			Title: &errMsg,
		}}
		return nil, errorObjs
	}
	included := tenantResourceObject.Included
	// TODO: what happen if departments exist? possible answer: this is verified and solved in the function CreateDepartmentsAndEmployeesBackend
	_, _, services := tenantUtils.ClassifyIncludedResourcesJSON(included) //this function returns services: "ServiceDB and servicesAPI"

	errMsg := tenantUtils.CheckTenantDatabaseArrayOfServiceAttributesJSON(services)
	if errMsg != "" {
		errCode := strconv.Itoa(http.StatusBadRequest)
		errorObjs := []didcommUtils.ErrorObject{{
			Code:  &errCode,
			Title: &errMsg,
		}}
		return nil, errorObjs
	}
	/*errMsg = tenantUtils.CheckTenantOrganizationAttributesJSON(tenantResourceObject.Attributes)
	if errMsg != "" {
		errCode := strconv.Itoa(http.StatusBadRequest)
		errorObjs := []didcommUtils.ErrorObject{{
			Code:  &errCode,
			Title: &errMsg,
		}}
		return nil, errorObjs
	}*/
	servicesResObjArray := ConvertArrayMapStrInterfaceToResourceObject(services)
	tenantBackend := tenantUtils.CreateByTenantResourceObject(tenantResourceObject)
	tenantBackend.AddManyServices(servicesResObjArray)
	edvHostProvider := storageHostUtils.GetHostProviderEDV()
	errMsg = storageHostUtils.SetTenantInHostEDV(edvHostProvider, tenantBackend)
	if errMsg != "" {
		errCode := strconv.Itoa(http.StatusInternalServerError)
		errorObjs := []didcommUtils.ErrorObject{{
			Code:  &errCode,
			Title: &errMsg,
		}}
		return nil, errorObjs
	}
	return tenantBackend.Services, errObjArray
}

func ConvertArrayMapStrInterfaceToResourceObject(services []map[string]interface{}) (resObjArray []didcommUtils.ResourceObject) {
	resourceBytes, _ := json.Marshal(services)
	_ = json.Unmarshal(resourceBytes, &resObjArray)
	return resObjArray
}

// Check if this function is really necessary as it only converts to a resource object and checks the types again
func PrepareConnectionDataToStore(servicesConnectionDataResourcesArray []didcommUtils.ResourceObject) (serviceConnectionData []didcommUtils.ResourceObject) {
	for _, singleServiceConn := range servicesConnectionDataResourcesArray {
		servType := singleServiceConn.Attributes["additionalType"]
		singleServiceAddType := tenantUtils.ConvertAdditionalTypeInServiceToString(servType)
		if strings.ToLower(singleServiceAddType) == "org.mongodb" || strings.ToLower(singleServiceAddType) == "identity-service" {
			serviceConnectionData = append(serviceConnectionData, singleServiceConn)
		}
	}
	return serviceConnectionData
}

func ConvertMapStrInterfaceToResourceObject(object map[string]interface{}) (resourceObj didcommUtils.ResourceObject) {
	var attsBytes = []byte{}
	attsBytes, _ = json.Marshal(object)
	_ = json.Unmarshal(attsBytes, &resourceObj)
	return resourceObj
}

var StoreServicesConnectionDataInTenantEDV = func(
	services []didcommUtils.ResourceObject,
	tenantAlternateName string,
) ([]didcommUtils.ResourceObject, []didcommUtils.ErrorObject) {
	resObjArray := []didcommUtils.ResourceObject{}
	errObjArray := []didcommUtils.ErrorObject{}

	// insert all connections in tenant
	connectionsResObjsArray, connectionsErrObjArray := edvConnection.SetManyServicesConnectionDataInEDV(services, tenantAlternateName)

	resObjArray = append(resObjArray, connectionsResObjsArray...)
	errObjArray = append(errObjArray, connectionsErrObjArray...)

	return resObjArray, errObjArray
}
