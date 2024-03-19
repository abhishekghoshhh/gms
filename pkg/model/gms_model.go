package model

type flowType string

const (
	TOKEN_FLOW             = flowType("TOKEN_FLOW")
	DEFAULT_FLOW           = flowType("DEFAULT_FLOW")
	CLIENT_CREDENTIAL_FLOW = flowType("CLIENT_CREDENTIAL_FLOW")
	PASSWORD_GRANT_FLOW    = flowType("PASSWORD_GRANT_FLOW")
)

type Flow interface {
	GetFlowType() flowType
}

type GmsModel struct {
	token             string
	subjectDn         string
	clientCert        string
	passwordGrantFlow bool
	groups            []string
}

func (gmsModel *GmsModel) GetFlowType() flowType {
	if gmsModel.token != "" {
		return TOKEN_FLOW
	} else if gmsModel.clientCert != "" && gmsModel.subjectDn != "" {
		if gmsModel.passwordGrantFlow {
			return PASSWORD_GRANT_FLOW
		} else {
			return CLIENT_CREDENTIAL_FLOW
		}
	} else {
		return DEFAULT_FLOW
	}
}

func GMS() *GmsModel {
	return &GmsModel{}
}
