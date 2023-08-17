package httputils

import (
	"io"
	"net/http"
	"strings"
	"bytes"

	decodeUtils "github.com/Universal-Health-Chain/common-utils-golang/decodeUtils"
	didcommUtils "github.com/Universal-Health-Chain/common-utils-golang/didCommunicationUtils"
	httpUtils "github.com/Universal-Health-Chain/common-utils-golang/httpUtils"
)

func DecodeHeadersAndBodyHTTP(r *http.Request) (decodedPayload *didcommUtils.DecodedRequestPayloadJAR, errMsg string) {
	// Get headers from the request
	openidHeaders := httpUtils.GetHttpHeaders(r)
	contentType := r.Header.Get("Content-Type")

	if strings.Contains(contentType, "jwt") || strings.Contains(contentType, "jwm") || strings.Contains(contentType, "didcomm") {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, "Failed to read request body."
		}

		// Ensure the body can be read again for any subsequent operations
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		compactJWT := string(bodyBytes)

		decodedPayload, errMsg := decodeUtils.GetDecodedRequestWithTokenData(compactJWT, *openidHeaders)
		if errMsg != "" {
			return nil, errMsg
		}

		return decodedPayload, ""
	}

	// ... handle other content types as needed
	return nil, "Unsupported content type."
}
