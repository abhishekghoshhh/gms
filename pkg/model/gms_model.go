package model

type GmsModel struct {
	token      string
	subjectDn  string
	clientCert string
	groups     []string
}

func (gmsModel *GmsModel) HasToken() bool {
	return gmsModel.token != ""
}

func (gmsModel *GmsModel) HasCert() bool {
	return gmsModel.clientCert != "" && gmsModel.subjectDn != ""
}

func GMS() *GmsModel {
	return &GmsModel{}
}
