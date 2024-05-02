package iam

import (
	"encoding/json"
	"testing"

	"github.com/abhishekghoshhh/gms/mocks"
	"github.com/abhishekghoshhh/gms/pkg/config"
	"github.com/abhishekghoshhh/gms/pkg/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestIamClient(t *testing.T) {
	t.Run("should fetch user id with given token", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHttpClient := mocks.NewMockHttpClient(ctrl)
		var iamConfig config.IamConfig

		iamClient := New(iamConfig, mockHttpClient)

		data := model.UserInfo{
			Userid:           "john",
			Name:             "John Doe",
			UserName:         "john",
			GivenName:        "John",
			FamilyName:       "Doe",
			Email:            "john@mail.com",
			OrganizationName: "abc",
		}

		jsonData, _ := json.Marshal(data)

		mockHttpClient.EXPECT().Send(gomock.Any()).Return(jsonData, nil)

		info, err := iamClient.FetchUserInfo("token")

		assert.NoError(t, err)
		assert.Equal(t, &data, info)
	})

	t.Run("should fetch user id with given token and userid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHttpClient := mocks.NewMockHttpClient(ctrl)
		iamConfig := config.IamConfig{
			Host: "host",
			Apis: config.IamApiConfig{
				FetchUserById: config.ApiConfig{
					Path:    "/",
					Timeout: 0,
				},
			},
		}

		iamClient := New(iamConfig, mockHttpClient)

		profileResponse := model.IamProfileResponse{
			Id:          "john",
			DisplayName: "john",
			Groups:      []model.Group{{Display: "group1"}},
			IndigoUserSchema: model.IndigoUserSchema{
				Certificates: []model.UserCertificate{
					{Primary: true,
						SubjectDn: "subject"},
				},
			},
		}
		jsonData, _ := json.Marshal(profileResponse)

		mockHttpClient.EXPECT().Send(gomock.Any()).Return(jsonData, nil)

		response, err := iamClient.FetchUserById("token", "userid")

		assert.NoError(t, err)
		assert.Equal(t, &profileResponse, response)
	})

	t.Run("should fetch client credential token", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHttpClient := mocks.NewMockHttpClient(ctrl)

		iamConfig := config.IamConfig{
			Host: "host",
			Apis: config.IamApiConfig{
				ClientCredentialToken: config.ApiConfig{
					Path:         "/token",
					Timeout:      0,
					ClientId:     "clientId",
					ClientSecret: "secret",
				},
			},
		}

		iamClient := New(iamConfig, mockHttpClient)

		token := model.Token{
			AccessToken: "acessToken",
			TokenType:   "a",
			Scope:       "scope",
			ExpiresIn:   0,
		}
		jsonData, _ := json.Marshal(token)

		mockHttpClient.EXPECT().Send(gomock.Any()).Return(jsonData, nil)

		credentialToken, err := iamClient.FetchClientCredentialToken()

		assert.NoError(t, err)
		assert.Equal(t, &token, credentialToken)
	})
}
