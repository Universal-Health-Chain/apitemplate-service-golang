package managers

import (
	didcommUtils "github.com/Universal-Health-Chain/common-utils-golang/didCommunicationUtils"
)

// In-memory store for message resource objects
var inMemoryStore = make(map[string][]didcommUtils.ResourceObject)

// Mocked function to store many message resource objects in memory.
func StoreManyResourceObjects(data *[]didcommUtils.ResourceObject, tenantAlternateName string, vaultId string) ([]didcommUtils.ResourceObject, *[]didcommUtils.ErrorObject) {
	key := tenantAlternateName + "_" + vaultId

	// Check if a slice already exists for this tenantAlternateName and vaultId
	if existingData, ok := inMemoryStore[key]; ok {
		// Append to the existing slice if it's already there
		existingData = append(existingData, *data...)
		inMemoryStore[key] = existingData
	} else {
		// Create a new entry in the map if it's not already there
		inMemoryStore[key] = *data
	}

	// Empty list for error objects
	var errorObjects *[]didcommUtils.ErrorObject

	// Return the stored data and an empty list for error objects
	return *data, errorObjects
}

// A utility function for testing purposes to retrieve stored data in memory
func GetStoredData(tenantAlternateName string, vaultId string) []didcommUtils.ResourceObject {
	key := tenantAlternateName + "_" + vaultId
	if data, ok := inMemoryStore[key]; ok {
		return data
	}
	return nil
}
