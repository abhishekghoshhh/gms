package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockCapabiltyBuilder struct {
}

func (*MockCapabiltyBuilder) Capabilities() string {
	return "mock-capability"
}

func TestGetTemplate(t *testing.T) {

	t.Run("should return capabilities in xml formal with success code", func(t *testing.T) {
		mockCapabilityBuilder := &MockCapabiltyBuilder{}

		capabilitiesApi := Capabilities(mockCapabilityBuilder)

		req := httptest.NewRequest("GET", "/template", nil)

		rr := httptest.NewRecorder()

		capabilitiesApi.GetTemplate(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expectedContentType := "application/xml"
		if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
			t.Errorf("handler returned wrong Content-Type header: got %v want %v", contentType, expectedContentType)
		}

		expectedBody := "mock-capability"
		if rr.Body.String() != expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
		}
	})
}
