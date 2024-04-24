package api

import (
	"net/http"

	"github.com/abhishekghoshhh/gms/pkg/iam"
	"github.com/labstack/echo"
)

type Handler struct {
	iamClient *iam.IamClient
}

func NewHandler(iamClient *iam.IamClient) *Handler {
	return &Handler{
		iamClient: iamClient,
	}
}

func (h *Handler) GetGroups(c echo.Context) error {
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

	return c.String(http.StatusOK, matchingGroups)
}
