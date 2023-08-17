package edvConnection

import (
	"github.com/Universal-Health-Chain/uhc-crypto-golang/edvUtils"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/storageHostUtils"
)

func ComposeQueryConnectionByAdditionalTypeInTenantEDV(attValue string, privateHostHmacKeyBytes []byte) edvUtils.Query {
	return storageHostUtils.ComposeGenericQuery(storageHostUtils.DepartmentKindIndexName, attValue, privateHostHmacKeyBytes)
}
