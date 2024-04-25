package api

import (
	"net/http"

	"github.com/abhishekghoshhh/gms/pkg/iam"
	"github.com/abhishekghoshhh/gms/pkg/logger"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type GroupsHandler struct {
	iamClient *iam.IamClient
}

func NewGetGroupsHandler(iamClient *iam.IamClient) *GroupsHandler {
	return &GroupsHandler{
		iamClient: iamClient,
	}
}

func (h *GroupsHandler) GetGroups(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	groups := c.Request().URL.Query()["group"]

	info, err := h.iamClient.FetchUserInfo(token)
	if err != nil {
		return err
	}

	clientCredentialToken, err := h.iamClient.FetchClientCredentialToken()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	iamProfile, err := h.iamClient.FetchUserById(clientCredentialToken.AccessToken, info.Userid)
	if err != nil {
		return err
	}
	logger.Debug("iam profile is", zap.Any("profile", iamProfile))
	matchingGroups := iamProfile.GetMatchingGroups(groups)

	c.Response().Header().Set("Content-Type", "text/plain")
	return c.String(http.StatusOK, matchingGroups)
}
