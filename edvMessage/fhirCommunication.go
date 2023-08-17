package edvMessage

//
//import (
//	"github.com/trustbloc/edv/pkg/restapi/models"
//)
//
//// Assuming these are declared somewhere else
//var GetTenantStorageHmacKeyBytes = func(tenantAlternateName string) []byte {
//	return []byte{}
//}
//var CreateIndexedHmacData = func(unprotectedIndexes []models.IndexedAttribute, hmacKeyBytes []byte) (protectedIndexData models.IndexedAttributeCollection) {
//	return models.IndexedAttributeCollection{}
//}
//
//// CreateFhirCommunicationIndexedHmacData creates indexed attribute collection for a FHIR Communication
//func CreateFhirCommunicationIndexedHmacData(fhirCommunication map[string]interface{}, tenantAlternateName string) models.IndexedAttributeCollection {
//	hmacKeyBytes := GetTenantStorageHmacKeyBytes(tenantAlternateName)
//	var unprotectedIndexes []models.IndexedAttribute
//
//	if subject, exists := fhirCommunication["subject"]; exists {
//		unprotectedIndexes = append(unprotectedIndexes, models.IndexedAttribute{Name: "subject", Value: subject.(string)})
//	}
//
//	if sender, exists := fhirCommunication["sender"]; exists {
//		unprotectedIndexes = append(unprotectedIndexes, models.IndexedAttribute{Name: "sender", Value: sender.(string)})
//	}
//
//	if recipients, exists := fhirCommunication["recipient"]; exists {
//		for i, recipient := range recipients.([]string) {
//			unprotectedIndexes = append(unprotectedIndexes, models.IndexedAttribute{Name: "recipient_" + string(i), Value: recipient})
//		}
//	}
//
//	return CreateIndexedHmacData(unprotectedIndexes, hmacKeyBytes)
//}
//
////endpoint
