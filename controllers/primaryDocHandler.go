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
		// Handle error as needed. Here's a basic approach:
		http.Error(w, "Failed to manage primary document", http.StatusInternalServerError)
		return
	}

	// Check Accept and Content-Type headers
	acceptHeader := r.Header.Get("Accept")
	contentTypeHeader := r.Header.Get("Content-Type")

	if strings.Contains(acceptHeader, "json") || strings.Contains(contentTypeHeader, "json") {
		// Return as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(primaryDoc)
		return
	} else if strings.Contains(acceptHeader, "jwm") || strings.Contains(acceptHeader, "jwt") || strings.Contains(acceptHeader, "didcomm") ||
		strings.Contains(contentTypeHeader, "jwm") || strings.Contains(contentTypeHeader, "jwt") || strings.Contains(contentTypeHeader, "didcomm") {

		// Convert primaryDoc to JWT and return as unsigned JWT
		jwtPayload := map[string]interface{}{
			"data": primaryDoc,
		}
		
		// Convert primaryDoc to a stringified version for JWT
		payloadBytes, _ := json.Marshal(jwtPayload)

		// Here we are using a utility function to create an unsigned JWT
		unsignedJWT := createTestUnsignedJWT(string(payloadBytes)) // Adjust as needed based on your utility function.

		w.Header().Set("Content-Type", "application/jwt")
		w.Write([]byte(unsignedJWT))
		return
	} else {
		// Default or handle other types
		http.Error(w, "Unsupported Accept or Content-Type", http.StatusBadRequest)
		return
	}
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