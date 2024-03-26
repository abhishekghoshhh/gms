package lib

import (
	"errors"
	"strings"

	"github.com/abhishekghoshhh/gms/pkg/client"
	"github.com/abhishekghoshhh/gms/pkg/model"
)

const (
	BEARER_TOKEN_PREFIX = "Bearer "
)

type AuthTokenFlow struct {
	iamClient *client.IamClient
}

func NewAuthTokenFlow(iamClient *client.IamClient) *AuthTokenFlow {
	return &AuthTokenFlow{
		iamClient,
	}
}

func (flow *AuthTokenFlow) GetGroups(gmsModel *model.GmsModel) (string, error) {
	token := gmsModel.Token()
	if !strings.HasPrefix(token, BEARER_TOKEN_PREFIX) {
		return "", errors.New("Invalid token " + token)
	}
	if iamProfile, err := flow.iamClient.FetchUser(token); err != nil {
		return "", err
	} else {
		return iamProfile.GetMatchingGroups(gmsModel.Groups()), nil
	}
}
