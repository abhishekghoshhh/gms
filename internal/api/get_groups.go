package api

import (
	"net/http"

	"github.com/abhishekghoshhh/gms/pkg/iam"
	"github.com/abhishekghoshhh/gms/pkg/logger"
	"github.com/labstack/echo"
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
		return err
	}

	iamProfile, err := h.iamClient.FetchUserById(clientCredentialToken.AccessToken, info.Userid)
	if err != nil {
		return err
	}

	matchingGroups := iamProfile.GetMatchingGroups(groups)

	logger.Info("matching groups " + matchingGroups)

	c.Response().Header().Set("Content-Type", "text/plain")
	return c.String(http.StatusOK, matchingGroups)
}
