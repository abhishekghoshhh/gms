package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abhishekghoshhh/gms/mocks"
	"go.uber.org/mock/gomock"
)

func TestGetGroups(t *testing.T) {

	t.Run("should return groups in text formal with success code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		gmsFlow := mocks.NewMockGmsFlow(ctrl)
		gmsFlow.EXPECT().GetGroups(gomock.Any()).Return("group1\n", nil)

		groupMembershipApi := GroupMembership(gmsFlow)

		req := httptest.NewRequest("GET", "/template", nil)

		rr := httptest.NewRecorder()

		groupMembershipApi.GetGroups(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expectedContentType := "text/plain"
		if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
			t.Errorf("handler returned wrong Content-Type header: got %v want %v", contentType, expectedContentType)
		}

		expectedBody := "group1\n"
		if rr.Body.String() != expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
		}
	})

	t.Run("should return error in text formal with bad request code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		gmsFlow := mocks.NewMockGmsFlow(ctrl)
		gmsFlow.EXPECT().GetGroups(gomock.Any()).Return("", errors.New("bad request error"))

		groupMembershipApi := GroupMembership(gmsFlow)

		req := httptest.NewRequest("GET", "/template", nil)

		rr := httptest.NewRecorder()

		groupMembershipApi.GetGroups(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expectedContentType := "text/plain"
		if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
			t.Errorf("handler returned wrong Content-Type header: got %v want %v", contentType, expectedContentType)
		}

		expectedBody := "bad request error"
		if rr.Body.String() != expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
		}
	})

}
