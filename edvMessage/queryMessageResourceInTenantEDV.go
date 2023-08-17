package edvMessage

import (
	"github.com/Universal-Health-Chain/uhc-crypto-golang/edvUtils"
	"github.com/Universal-Health-Chain/uhc-crypto-golang/storageHostUtils"
)

func ComposeQueryMessageResourceByAdditionalTypeInTenantEDV(attValue string, privateHostHmacKeyBytes []byte) edvUtils.Query {
	return storageHostUtils.ComposeGenericQuery(storageHostUtils.MessageResourceKindIndexName, attValue, privateHostHmacKeyBytes)
}
