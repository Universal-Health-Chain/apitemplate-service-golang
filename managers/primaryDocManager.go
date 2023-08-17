package managers

import (
	httpUtils "apitemplate-service-golang/utils/httpUtils"
	"net/http"

	didcommUtils "github.com/Universal-Health-Chain/common-utils-golang/didCommunicationUtils"
	"github.com/go-chi/chi/v5"
)

func ManagePostPrimaryDocToVaultEDV(r *http.Request) (didcommUtils.PrimaryDocument, error) {
	tenantAlternateName := chi.URLParam(r, "tenantAlternateName")
	vaultId := chi.URLParam(r, "vaultId")

	decodedRequest, errMsg := httpUtils.DecodeHeadersAndBodyHTTP(r)
	if errMsg != "" {
		return didcommUtils.PrimaryDocument{
			Errors: &[]didcommUtils.ErrorObject{
				{
					Detail: &errMsg,
				},
			},
		}, nil
	}

	// Directly using decodedRequest.Body, since it's already of type didcommUtils.PrimaryDocument
	primaryDoc := decodedRequest.Body
	storedResourceObjects, errorObjects := StoreManyResourceObjects(&primaryDoc.Data, tenantAlternateName, vaultId)

	return didcommUtils.PrimaryDocument{
		Data:   storedResourceObjects,
		Errors: errorObjects,
	}, nil
}
