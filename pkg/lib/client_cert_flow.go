package lib

import (
	"errors"
	"strings"
	"sync"

	"github.com/abhishekghoshhh/gms/pkg/client"
	"github.com/abhishekghoshhh/gms/pkg/logger"
	"github.com/abhishekghoshhh/gms/pkg/model"
	"github.com/abhishekghoshhh/gms/pkg/util"
	"go.uber.org/zap"
)

type ClientCertFlow struct {
	iamClient                 *client.IamClient
	isPasswordGrantFlowActive bool
	passwordGrantConfig       *model.PasswordGrantConfig
	clientCredentialConfig    *model.ClientCredentialConfig
}

func NewClientCertFlow(iamClient *client.IamClient, isPasswordGrantFlowActive string,
	passwordGrantConfig *model.PasswordGrantConfig, clientCredentialConfig *model.ClientCredentialConfig) (*ClientCertFlow, error) {

	passwordGrantFlowEnabled := util.Bool(isPasswordGrantFlowActive)
	if passwordGrantFlowEnabled && !passwordGrantConfig.IsValid() {
		return nil, errors.New("password grant flow is active but password grant config is not valid")
	} else if !passwordGrantFlowEnabled && !clientCredentialConfig.IsValid() {
		return nil, errors.New("password grant flow is inactive but client credential config is not valid")
	}

	return &ClientCertFlow{
		iamClient:                 iamClient,
		isPasswordGrantFlowActive: passwordGrantFlowEnabled,
		passwordGrantConfig:       passwordGrantConfig,
		clientCredentialConfig:    clientCredentialConfig,
	}, nil
}

func (flow *ClientCertFlow) GetGroups(data *model.GmsModel) (string, error) {
	if flow.isPasswordGrantFlowActive {
		return flow.fromPasswordGrant(data)
	} else {
		return flow.fromClientCredential(data)
	}
}

func (flow *ClientCertFlow) fromPasswordGrant(data *model.GmsModel) (string, error) {
	logger.Info("Password grant flow for GMS group search is executing")

	clientId := flow.passwordGrantConfig.ClientId()
	clientSecret := flow.passwordGrantConfig.ClientSecret()
	adminName := flow.passwordGrantConfig.UserName()
	adminPassword := flow.passwordGrantConfig.Password()

	accessToken, err := flow.iamClient.FetchAccessTokenForPasswordGrantFlow(adminName, adminPassword, clientId, clientSecret)
	if err != nil {
		logger.Error("failed to fetch access token " + err.Error())
		return "", errors.New("failed to fetch the accessToken")
	}

	userListResponse, err := flow.iamClient.FetchUserByCertSubject(accessToken.AccessToken, data.Subject())
	if err != nil {
		return "", errors.New("failed to fetch userList")
	}

	if len(userListResponse.Resources) == 0 {
		return "", errors.New("current client certificate is not linked with any user")
	}

	return userListResponse.Resources[0].GetMatchingGroups(data.Groups()), nil
}

func (flow *ClientCertFlow) fromClientCredential(data *model.GmsModel) (string, error) {
	logger.Info("Client credential flow for GMS group search is executing")

	clientId := flow.clientCredentialConfig.ClientId()
	clientSecret := flow.clientCredentialConfig.ClientSecret()

	accessToken, err := flow.iamClient.FetchAccessTokenForClientCredentialFlow(clientId, clientSecret)
	if err != nil {
		logger.Error("failed to fetch access token " + err.Error())
		return "", errors.New("failed to fetch the accessToken")
	}
	profileListResponse, err := flow.iamClient.FetchUserCount(accessToken.AccessToken)
	if err != nil {
		logger.Error("unable to fetch user profile list " + err.Error())
		return "", errors.New("unable to fetch user profile list")
	}
	logger.Debug("", zap.Int("user count is", profileListResponse.TotalResults))

	batch := model.NewBatch(
		profileListResponse.TotalResults,
		flow.clientCredentialConfig.BatchSize(),
		flow.clientCredentialConfig.MaxBatchCount(),
	)

	subjectDn := data.Subject()
	clientCert := strings.Replace(data.ClientCert(), "\t", "\n", -1)

	var wg sync.WaitGroup
	responseChannel := make(chan *model.IamProfileResponse, batch.Count)

	for _, batchStartIdx := range batch.Indices {
		wg.Add(1)

		go func(batchStart int) {
			defer wg.Done()

			profileList, err := flow.iamClient.FetchUsersInBatch(
				accessToken.AccessToken,
				batchStart,
				batch.Size,
			)

			if err != nil {
				logger.Error("internal error " + err.Error())
				responseChannel <- nil
			}
			for _, profile := range profileList.Resources {
				if profile.HasMatchingCert(subjectDn, clientCert) {
					logger.Info("Found the user with client cert")
					responseChannel <- &profile
				}
			}
			responseChannel <- nil

		}(batchStartIdx)
	}
	wg.Wait()
	close(responseChannel)

	for profile := range responseChannel {
		if profile != nil {
			return profile.GetMatchingGroups(data.Groups()), nil
		}
	}

	return "", errors.New("current client certificate is not linked with any user")
}
