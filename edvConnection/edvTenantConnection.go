package edvConnection

import (
	"encoding/json"
	"fmt"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/didcommUtils"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/edvUtils"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/storageHostUtils"
	"net/http"
	"strconv"
	"strings"
)

func createServiceConnectionDataIndexedHmacData(service didcommUtils.ResourceObject, tenantAlternateName string) (protectedIndexedData edvUtils.IndexedData) {

	// TODO: get the tenantHmacKeyBytes
	tenantHmacKeyBytes := storageHostUtils.GetTenantStorageHmacKeyBytes(tenantAlternateName) // TODO
	resourceBytes, _ := json.Marshal(service)
	plaintext := string(resourceBytes)
	unprotectedIndexes := []edvUtils.IndexedAttribute{
		{ // use this one to find a concrete department
			Name: service.Type, //"ServiceAPI",
			//Value:  *departmentBackend.Attributes.Main.AdditionalType, // lowercase? e.g.: "type:ict" or "org.hl7.fhir.valueset.service-type:51"
			Value:  plaintext,
			Unique: true,
		},
	}

	return edvUtils.CreateIndexedHmacData(unprotectedIndexes, tenantHmacKeyBytes)
}

var SetServiceConnectionDataInTenantEDV = func(service didcommUtils.ResourceObject, tenantAlternateName string) (errMsg string) {
	tenantVault := ""
	switch strings.ToLower(service.Type) {
	case "servicedb":
		tenantVault = storageHostUtils.TenantVaultDepartments
	case "serviceapi":
		tenantVault = storageHostUtils.TenantVaultConnection
	default:
		return "cannot set service connection in tenant. Improper service type"
	}
	if service.Type == "" {
		return "cannot set service connection in tenant"
	}
	resourceBytes, _ := json.Marshal(service)
	plaintext := string(resourceBytes)
	jwe, errMsg := storageHostUtils.EncryptHostDataEDV(&plaintext)
	if errMsg != "" {
		return errMsg
	}

	hmacProtectedIndexData := createServiceConnectionDataIndexedHmacData(service, tenantAlternateName)

	// TODO: get the tenantHmacKeyBytes
	tenantHmacKeyBytes := storageHostUtils.GetTenantStorageHmacKeyBytes(tenantAlternateName) // TODO

	encryptedDocument := edvUtils.EncryptedDocument{
		// TODO: add the GetUID() method to the backend resource and use it instead of the DID?
		ID:          edvUtils.GenerateIdentifierDeterministicallyForEDV(service.GetDID(), tenantHmacKeyBytes), // truncated to 16 characters in base58 encoding
		IndexedData: []edvUtils.IndexedData{hmacProtectedIndexData},
		JWE:         jwe, // it can be unmarshalled to JWEncryptionGo
	}

	// TODO:
	tenantProviderEDV := storageHostUtils.GetTenantProviderEDV(tenantAlternateName)
	err := tenantProviderEDV.Put(tenantVault, encryptedDocument)
	if err != nil {
		return fmt.Sprint(err.Error())
	} else {
		return ""
	}
}

// TODO: only one result shall be received, so manage the error if len > 1
// TODO: tenantPrivateHmacKey
/*var GetConnectionsByAdditionalTypeInTenantEDV = func(tenantAlternateName, additionalType string) (connectionData *models.ConnectionData, errMsg string) {
	tenantProviderEDV := storageHostUtils.GetTenantProviderEDV(tenantAlternateName)
	query := ComposeQueryConnectionByAdditionalTypeInTenantEDV(additionalType, storageHostUtils.GetTenantStorageHmacKeyBytes(tenantAlternateName))
	encryptedResults, err := tenantProviderEDV.Query(storageHostUtils.TenantVaultConnection, query)
	if err != nil {
		return nil, errMsg
	}

	fmt.Printf("Number of tenants found by alternate name in the Host's EDV = %v", len(encryptedResults))

	if len(encryptedResults) != 1 {
		return nil, "error finding document"
	}

	edvDocument := encryptedResults[0]
	decryptedDataBytes, _ := storageHostUtils.DecryptHostDataEDV(&edvDocument.JWE)

	connectionData = &models.ConnectionData{} // non empty object to avoid unmarshall errors
	err = json.Unmarshal(decryptedDataBytes, &connectionData)
	if err != nil {
		return nil, errMsg
	}

	return connectionData, ""
}*/

var SetManyServicesConnectionDataInEDV = func(services []didcommUtils.ResourceObject, tenantAlternateName string) ([]didcommUtils.ResourceObject, []didcommUtils.ErrorObject) {
	if services == nil {
		errCode := strconv.Itoa(http.StatusBadRequest)
		errTitle := "Bad request of services insertion: no data"
		errorObjs := []didcommUtils.ErrorObject{{
			Code:  &errCode,
			Title: &errTitle,
		}}
		/* return "cannot add departments, cannot be empty" */
		return []didcommUtils.ResourceObject{}, errorObjs
	}

	successfulServiceConnectionDataArray := []didcommUtils.ResourceObject{}
	failingServiceConnectionDataArray := []didcommUtils.ResourceObject{}

	for _, singleServiceConnectionData := range services {
		errMsg := SetServiceConnectionDataInTenantEDV(singleServiceConnectionData, tenantAlternateName)
		if errMsg != "" {
			fmt.Println(errMsg)
			failingServiceConnectionDataArray = append(failingServiceConnectionDataArray, singleServiceConnectionData)
		} else {
			successfulServiceConnectionDataArray = append(successfulServiceConnectionDataArray, singleServiceConnectionData)
		}
	}

	failErrObjs := []didcommUtils.ErrorObject{}

	if len(failingServiceConnectionDataArray) > 0 {
		errCode := fmt.Sprint(http.StatusInternalServerError)
		errMsg := "Failed to write in DB"
		failErrObjs = *(ConvertServiceFailingConnectionDataToErrorObjects(failingServiceConnectionDataArray, errCode, errMsg))
	}

	successResourceObjectArray := successfulServiceConnectionDataArray

	return successResourceObjectArray, failErrObjs
}

var ConvertServiceFailingConnectionDataToErrorObjects = func(existingServiceConnections []didcommUtils.ResourceObject, errCode, errDetail string) *[]didcommUtils.ErrorObject {
	connectionDataErrorArray := []didcommUtils.ErrorObject{}
	for _, existingServiceConnection := range existingServiceConnections {
		id := existingServiceConnection.GetDID()
		departmentError := didcommUtils.ErrorObject{
			ID:     &id,
			Code:   &errCode,
			Detail: &errDetail,
		}
		connectionDataErrorArray = append(connectionDataErrorArray, departmentError)
	}
	return &connectionDataErrorArray
}
