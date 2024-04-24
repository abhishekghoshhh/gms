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

	logger.Info("token is", zap.Any("token", token))
	logger.Info("group is", zap.Any("info", groups))

	info, err := h.iamClient.FetchUserInfo(token)
	if err != nil {
		return err
	}

	logger.Info("info is", zap.Any("info", info))

	clientCredentialToken, err := h.iamClient.FetchClientCredentialToken()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("clientCredentialToken is", zap.Any("profile", clientCredentialToken))

	iamProfile, err := h.iamClient.FetchUserById(clientCredentialToken.AccessToken, info.Userid)
	if err != nil {
		return err
	}

	logger.Info("iam profile is", zap.Any("profile", iamProfile))

	matchingGroups := iamProfile.GetMatchingGroups(groups)

	logger.Info("matching groups " + matchingGroups)

	c.Response().Header().Set("Content-Type", "text/plain")
	return c.String(http.StatusOK, matchingGroups)
}
