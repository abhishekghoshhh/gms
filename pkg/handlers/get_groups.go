package handlers

import (
	"github.com/abhishekghoshhh/gms/pkg/iamclient"
	"github.com/abhishekghoshhh/gms/pkg/logger"
	"net/http"
)

type Handler struct {
	iamClient *iamclient.IamClient
}

func NewHandler(client iamclient.IamClient) *Handler {
	return &Handler{
		iamClient: &client,
	}
}

func (h *Handler) GetGroups(responseWriter http.ResponseWriter, request *http.Request) {
	token := request.Header.Get("Authorization")
	groups := request.URL.Query()["group"]

	getGroups, err := h.getGroups(token, groups)
	if err != nil {
		 responseWriter.WriteHeader(400)
	}
	_, err = responseWriter.Write([]byte(getGroups))
	if err != nil {
		logger.Error("Error in response writer")
	}
}

func (h *Handler) getGroups(token string, groups []string) (string, error) {

	info, err := h.iamClient.FetchUserInfo(token)
	if err != nil {
		return "", err
	}
	clientCredentialToken, err := h.iamClient.FetchClientCredentialToken()
	if err != nil {
		return "", err
	}

	resp, err := h.iamClient.FetchUserById(clientCredentialToken.AccessToken, info.Userid)
	if err != nil {
		return "", err
	}

	matchingGroups := resp.GetMatchingGroups(groups)

	if err != nil {
		return "", err
	}

	return matchingGroups, nil
}

