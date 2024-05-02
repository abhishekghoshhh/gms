package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abhishekghoshhh/gms/mocks"
	"github.com/abhishekghoshhh/gms/pkg/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	userToken   = "Bearer token"
	userId      = "1234"
	clientToken = "Bearer client token"

	defaultUserInfo    = &model.UserInfo{Userid: userId}
	defaultClientToken = &model.Token{AccessToken: clientToken}
	defaultIamProfile  = &model.IamProfileResponse{
		Groups: []model.Group{
			{Display: "group1"},
			{Display: "group2"},
			{Display: "group3"},
		},
	}

	internalServerError = echo.ErrInternalServerError
)

func TestGetGroups(t *testing.T) {
	t.Run("should return groups in text formal with success code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		iamClient := mocks.NewMockClient(ctrl)
		handler := NewGetGroupsHandler(iamClient)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/gms/search", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Request().Header.Set("Authorization", userToken)
		c.Request().URL.RawQuery = "group=group1&group=group2"

		iamClient.EXPECT().FetchUserInfo(gomock.Any()).Return(defaultUserInfo, nil)
		iamClient.EXPECT().FetchClientCredentialToken().Return(defaultClientToken, nil)
		iamClient.EXPECT().FetchUserById(clientToken, userId).Return(defaultIamProfile, nil)

		err := handler.GetGroups(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "text/plain", rec.Header().Get("Content-Type"))
		assert.Equal(t, "group1\ngroup2\n", rec.Body.String())
	})

	t.Run("should return all groups in text on empty group query formal with success code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		iamClient := mocks.NewMockClient(ctrl)
		handler := NewGetGroupsHandler(iamClient)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/gms/search", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Request().Header.Set("Authorization", userToken)

		iamClient.EXPECT().FetchUserInfo(gomock.Any()).Return(defaultUserInfo, nil)
		iamClient.EXPECT().FetchClientCredentialToken().Return(defaultClientToken, nil)
		iamClient.EXPECT().FetchUserById(clientToken, userId).Return(defaultIamProfile, nil)

		err := handler.GetGroups(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "text/plain", rec.Header().Get("Content-Type"))
		assert.Equal(t, "group1\ngroup2\ngroup3\n", rec.Body.String())
	})

	t.Run("should return bad request on user token missing ", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		iamClient := mocks.NewMockClient(ctrl)
		handler := NewGetGroupsHandler(iamClient)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/gms/search", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.GetGroups(c).(*echo.HTTPError)

		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, err.Code)
	})

	t.Run("should return internal server error if userinfo call fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		iamClient := mocks.NewMockClient(ctrl)
		handler := NewGetGroupsHandler(iamClient)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/gms/search", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Request().Header.Set("Authorization", userToken)

		iamClient.EXPECT().FetchUserInfo(gomock.Any()).Return(nil, internalServerError)

		err := handler.GetGroups(c).(*echo.HTTPError)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.Code)
	})

	t.Run("should return internal server error if client credential call fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		iamClient := mocks.NewMockClient(ctrl)
		handler := NewGetGroupsHandler(iamClient)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/gms/search", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Request().Header.Set("Authorization", userToken)

		iamClient.EXPECT().FetchUserInfo(gomock.Any()).Return(defaultUserInfo, nil)
		iamClient.EXPECT().FetchClientCredentialToken().Return(nil, internalServerError)

		err := handler.GetGroups(c).(*echo.HTTPError)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.Code)
	})

	t.Run("should return internal server error if fetch by user id call fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		iamClient := mocks.NewMockClient(ctrl)
		handler := NewGetGroupsHandler(iamClient)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/gms/search", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Request().Header.Set("Authorization", userToken)

		iamClient.EXPECT().FetchUserInfo(gomock.Any()).Return(defaultUserInfo, nil)
		iamClient.EXPECT().FetchClientCredentialToken().Return(defaultClientToken, nil)
		iamClient.EXPECT().FetchUserById(clientToken, userId).Return(nil, internalServerError)

		err := handler.GetGroups(c).(*echo.HTTPError)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.Code)
	})
}
