package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"apitemplate-service-golang/managers"

	didcommUtils "github.com/Universal-Health-Chain/common-utils-golang/didCommunicationUtils"
)

// A mock manager function for testing purposes
func mockManagePostPrimaryDocToVaultEDV(r *http.Request) (didcommUtils.PrimaryDocument, error) {
	// Mock logic to return a PrimaryDocument
	var resourceObjectHashedId = "test-id"

	return didcommUtils.PrimaryDocument{
		/*... Some mock data ...*/
		Data: []didcommUtils.ResourceObject{
			{IdHashed: resourceObjectHashedId, Type: "test-type"},
		},
	}, nil
}

// Using table-driven tests as an effective method to include assertions for response body and code
func TestHandlePostPrimaryDocToStore_ContentTypes(t *testing.T) {
	// Mocking the manager function
	managers.ManagePostPrimaryDocToVaultEDV = mockManagePostPrimaryDocToVaultEDV

	for _, contentType := range supportedContentTypes {
		// Test for supported Content-Type
		req := httptest.NewRequest("POST", "http://example.com/foo", nil)
		req.Header.Set("Content-Type", contentType)
		w := httptest.NewRecorder()

		HandlePostPrimaryDocToStore(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK for content type %s, but got %d", contentType, resp.StatusCode)
		}
	}

	// Test for unsupported Content-Type
	req := httptest.NewRequest("POST", "http://example.com/foo", nil)
	req.Header.Set("Content-Type", "unsupported/type")
	w := httptest.NewRecorder()

	HandlePostPrimaryDocToStore(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest for unsupported content type, but got %d", resp.StatusCode)
	}

	// Test for empty Content-Type
	reqEmptyContentType := httptest.NewRequest("POST", "http://example.com/foo", nil)
	wEmptyContentType := httptest.NewRecorder()

	HandlePostPrimaryDocToStore(wEmptyContentType, reqEmptyContentType)

	respEmptyContentType := wEmptyContentType.Result()
	if respEmptyContentType.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest for empty content type, but got %d", respEmptyContentType.StatusCode)
	}
}

func TestHandlePostPrimaryDocToStore(t *testing.T) {
	// Mocking the manager function
	managers.ManagePostPrimaryDocToVaultEDV = mockManagePostPrimaryDocToVaultEDV

	tests := []struct {
		name           string
		contentType    string
		expectedOutput string // This might be a constant or a representative slice of the output.
		expectedCode   int
	}{
		{
			name:           "Content-Type: json",
			contentType:    "json",
			expectedOutput: "{\"...some json...\"}", // Assuming your mock manager returns this.
			expectedCode:   http.StatusOK,
		},
		{
			name:           "Content-Type: jwm",
			contentType:    "jwm",
			expectedOutput: "header.payload.", // As per createTestUnsignedJWT function.
			expectedCode:   http.StatusOK,
		},
		// ... (repeat for each content type)
		{
			name:           "Content-Type: unsupported/type",
			contentType:    "unsupported/type",
			expectedOutput: "", // No output for unsupported content type.
			expectedCode:   http.StatusBadRequest,
		},
		{
			name:           "Content-Type: empty",
			contentType:    "",
			expectedOutput: "", // No output for empty content type.
			expectedCode:   http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "http://example.com/foo", nil)
			req.Header.Set("Content-Type", tt.contentType)
			rr := httptest.NewRecorder()

			HandlePostPrimaryDocToStore(rr, req)

			assert.Equal(t, tt.expectedOutput, rr.Body.String())
			assert.Equal(t, tt.expectedCode, rr.Code)
		})
	}
}
