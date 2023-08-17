package edvMessage

import (
	"apitemplate-service-golang/tenantConnection"
	"encoding/json"
	"fmt"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/edvUtils"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/storageHostUtils"
	"net/http"
	"strconv"

	"github.com/Universal-Health-Chain/uhc-crypto-golang/didcommUtils"
)

// TODO: get the tenantHmacKeyBytes
// TODO: lowercase the values before generating HMAC?
func createMessageResourceIndexedHmacData(messageResource didcommUtils.ResourceObject, tenantAlternateName string) (protectedIndexedData edvUtils.IndexedData) {

	// TODO: get the tenantHmacKeyBytes
	tenantHmacKeyBytes := storageHostUtils.GetTenantStorageHmacKeyBytes(tenantAlternateName) // TODO

	unprotectedIndexes := []edvUtils.IndexedAttribute{
		{ // use this one to find a concrete department
			Name:   "additionalType",
			Value:  messageResource.Type, // lowercase? e.g.: "type:ict" or "org.hl7.fhir.valueset.service-type:51"
			Unique: true,
		},
	}

	return edvUtils.CreateIndexedHmacData(unprotectedIndexes, tenantHmacKeyBytes)
}

var setMessageResourceInTenantEDV = func(messageResource didcommUtils.ResourceObject, tenantAlternateName string) (errMsg string) {
	if messageResource.Type != "message" { //check if this name corresponds
		return "cannot set message resource in tenant"
	}
	var plaintext string
	plaintext = Stringify(messageResource.Attachments)

	jwe, errMsg := storageHostUtils.EncryptHostDataEDV(&plaintext)
	if errMsg != "" {
		return errMsg
	}

	hmacProtectedIndexData := createMessageResourceIndexedHmacData(messageResource, tenantAlternateName)

	// TODO: get the tenantHmacKeyBytes
	tenantHmacKeyBytes := storageHostUtils.GetTenantStorageHmacKeyBytes(tenantAlternateName) // TODO

	encryptedDocument := edvUtils.EncryptedDocument{
		// TODO: add the GetUID() method to the backend resource and use it instead of the DID?
		ID:          edvUtils.GenerateIdentifierDeterministicallyForEDV(messageResource.GetDID(), tenantHmacKeyBytes), // truncated to 16 characters in base58 encoding
		IndexedData: []edvUtils.IndexedData{hmacProtectedIndexData},
		JWE:         jwe, // it can be unmarshalled to JWEncryptionGo
	}

	// TODO:
	tenantProviderEDV := storageHostUtils.GetTenantProviderEDV(tenantAlternateName)
	err := tenantProviderEDV.Put(storageHostUtils.TenantVaultDepartments, encryptedDocument)
	if err != nil {
		return fmt.Sprint(err.Error())
	} else {
		return ""
	}
}

// TODO: only one result shall be received, so manage the error if len > 1
// TODO: tenantPrivateHmacKey
var GetMessageResourceByAdditionalTypeInTenantEDV = func(tenantAlternateName, additionalType string) (messageResource *didcommUtils.ResourceObject, errMsg string) {
	tenantProviderEDV := storageHostUtils.GetTenantProviderEDV(tenantAlternateName)
	query := ComposeQueryMessageResourceByAdditionalTypeInTenantEDV(additionalType, storageHostUtils.GetTenantStorageHmacKeyBytes(tenantAlternateName))
	encryptedResults, err := tenantProviderEDV.Query(storageHostUtils.TenantVaultDepartments, query)
	if err != nil {
		return nil, errMsg
	}

	fmt.Printf("Number of tenants found by alternate name in the Host's EDV = %v", len(encryptedResults))

	if len(encryptedResults) != 1 {
		return nil, "error finding document"
	}

	edvDocument := encryptedResults[0]
	decryptedDataBytes, _ := storageHostUtils.DecryptHostDataEDV(&edvDocument.JWE)

	messageResource = &didcommUtils.ResourceObject{} // non-empty object to avoid unmarshall errors
	err = json.Unmarshal(decryptedDataBytes, &messageResource)
	if err != nil {
		return nil, errMsg
	}

	return messageResource, ""
}

//const collectionName = config.CollDepartments

// TODO: separate InsertBatch from InsertTx
var SetManyMessageResourcesInEDV = func(messageResourceSlice []map[string]interface{}, tenantAlternateName string) ([]didcommUtils.ResourceObject, []didcommUtils.ErrorObject) {
	if messageResourceSlice == nil {
		errCode := strconv.Itoa(http.StatusBadRequest)
		errTitle := "Bad request: no data"
		errorObjs := []didcommUtils.ErrorObject{{
			Code:  &errCode,
			Title: &errTitle,
		}}

		/* return "cannot add departments, cannot be empty" */
		return []didcommUtils.ResourceObject{}, errorObjs
	}

	messageResource := tenantConnection.ConvertArrayMapStrInterfaceToResourceObject(messageResourceSlice)

	successfulMessageResourceArray := []didcommUtils.ResourceObject{}
	failingMessageResourceArray := []didcommUtils.ResourceObject{}
	var singleObject didcommUtils.ResourceObject

	for _, singleObject = range messageResource {
		errMsg := setMessageResourceInTenantEDV(singleObject, tenantAlternateName)
		if errMsg != "" {
			fmt.Println(errMsg)
			failingMessageResourceArray = append(failingMessageResourceArray, singleObject)
		} else {
			successfulMessageResourceArray = append(successfulMessageResourceArray, singleObject)
		}
	}

	failErrObjs := []didcommUtils.ErrorObject{}

	if len(failingMessageResourceArray) > 0 {
		errCode := fmt.Sprint(http.StatusInternalServerError)
		errMsg := "Could not write in DB"
		failErrObjs = *ConvertFailingMessageResourcesToErrorObjects(failingMessageResourceArray, errCode, errMsg)
	}

	successResourceObjectArray := successfulMessageResourceArray

	return successResourceObjectArray, failErrObjs
}

var ConvertFailingMessageResourcesToErrorObjects = func(existingMessages []didcommUtils.ResourceObject, errCode, errDetail string) *[]didcommUtils.ErrorObject {
	connectionDataErrorArray := []didcommUtils.ErrorObject{}
	for _, singleExistingMessage := range existingMessages {
		id := singleExistingMessage.GetDID()
		departmentError := didcommUtils.ErrorObject{
			ID:     &id,
			Code:   &errCode,
			Detail: &errDetail,
		}
		connectionDataErrorArray = append(connectionDataErrorArray, departmentError)
	}
	return &connectionDataErrorArray
}

func Stringify(obj interface{}) string {
	resourceBytes, _ := json.Marshal(obj)
	return string(resourceBytes)
}
