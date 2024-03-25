package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abhishekghoshhh/gms/mocks"
	"go.uber.org/mock/gomock"
)

func TestGetTemplate(t *testing.T) {

	t.Run("should return capabilities in xml formal with success code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCapabilityBuilder := mocks.NewMockCapabilityBuilder(ctrl)
		mockCapabilityBuilder.EXPECT().Capabilities().Return("gomock-capability")

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

		expectedBody := "gomock-capability"
		if rr.Body.String() != expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
		}
	})
}
