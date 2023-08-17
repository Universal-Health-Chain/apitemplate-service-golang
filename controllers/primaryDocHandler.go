package controllers

import (
	managers "apitemplate-service-golang/managers" // replace with the actual path
	"encoding/base64"
	"encoding/json"
	"net/http"

	"strings"
)

func HandlePostPrimaryDocToStore(w http.ResponseWriter, r *http.Request) {
	primaryDoc, err := managers.ManagePostPrimaryDocToVaultEDV(r)

	if err != nil {
		http.Error(w, "Failed to manage primary document", http.StatusInternalServerError)
		return
	}

	// Check Accept and Content-Type headers
	acceptHeader := r.Header.Get("Accept")
	contentTypeHeader := r.Header.Get("Content-Type")

	// Using a helper function for readability and to avoid repetitive code
	if isSupportedContentType(contentTypeHeader) {
		if strings.Contains(acceptHeader, "json") || strings.Contains(contentTypeHeader, "json") {
			// Return as JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(primaryDoc)
		} else if strings.Contains(acceptHeader, "jwm") || strings.Contains(acceptHeader, "jwt") || strings.Contains(acceptHeader, "didcomm") ||
			strings.Contains(contentTypeHeader, "jwm") || strings.Contains(contentTypeHeader, "jwt") || strings.Contains(contentTypeHeader, "didcomm") {

			// Convert primaryDoc to JWT and return as unsigned JWT
			jwtPayload := map[string]interface{}{
				"data": primaryDoc,
			}

			// Convert primaryDoc to a stringified version for JWT
			payloadBytes, _ := json.Marshal(jwtPayload)

			// Here we are using a utility function to create an unsigned JWT
			unsignedJWT := createTestUnsignedJWT(string(payloadBytes))

			w.Header().Set("Content-Type", "application/jwt")
			w.Write([]byte(unsignedJWT))
		} else {
			http.Error(w, "Unsupported Accept header", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "Unsupported Content-Type", http.StatusBadRequest)
		return
	}
}

// Helper function to check for supported content types
func isSupportedContentType(contentType string) bool {
	for _, t := range supportedContentTypes {
		if strings.Contains(contentType, t) {
			return true
		}
	}
	return false
}

func createTestUnsignedJWT(stringifiedPayload string) string {
	headers := map[string]interface{}{
		"alg": "none",
	}

	// Encode headers
	headerBytes, _ := json.Marshal(headers)
	encodedHeader := base64.RawURLEncoding.EncodeToString(headerBytes)

	// Encode payload
	encodedPayload := base64.RawURLEncoding.EncodeToString([]byte(stringifiedPayload))

	// JWT is 'header.payload.'
	return encodedHeader + "." + encodedPayload + "."
}
