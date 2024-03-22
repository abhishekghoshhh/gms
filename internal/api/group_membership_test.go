package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abhishekghoshhh/gms/pkg/model"
)

type MockGmsFlow struct {
	getGroupHandle func(gmsModel *model.GmsModel) (string, error)
}

func (gmsFlow *MockGmsFlow) GetGroups(gmsModel *model.GmsModel) (string, error) {
	return gmsFlow.getGroupHandle(gmsModel)
}

func TestGetGroupsWithSuccess(t *testing.T) {
	mockGmsFlow := &MockGmsFlow{
		getGroupHandle: func(gmsModel *model.GmsModel) (string, error) {
			return "group1\n", nil
		},
	}

	groupMembershipApi := GroupMembership(mockGmsFlow)

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
}

func TestGetGroupsWithError(t *testing.T) {

	mockGmsFlow := &MockGmsFlow{
		getGroupHandle: func(gmsModel *model.GmsModel) (string, error) {
			return "", errors.New("bad request error")
		},
	}

	groupMembershipApi := GroupMembership(mockGmsFlow)

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
}
