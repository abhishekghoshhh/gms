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
func (gmsModel *GmsModel) Token() string {
	return gmsModel.token
}

func (gmsModel *GmsModel) HasCert() bool {
	return gmsModel.clientCert != "" && gmsModel.subjectDn != ""
}

func (gmsModel *GmsModel) HasGroups() bool {
	return len(gmsModel.groups) != 0
}

func (gmsModel *GmsModel) Groups() []string {
	return gmsModel.groups
}

func (gmsModel *GmsModel) Subject() string { return gmsModel.subjectDn }

func GMS(token, subjectDn, clientCert string, groups []string) *GmsModel {
	return &GmsModel{
		token,
		subjectDn,
		clientCert,
		groups,
	}
}
