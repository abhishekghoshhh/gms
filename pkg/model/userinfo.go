package model

type UserInfo struct {
	Userid           string `json:"sub"`
	Name             string `json:"name"`
	UserName         string `json:"preferred_username"`
	GivenName        string `json:"given_name"`
	FamilyName       string `json:"family_name"`
	Email            string `json:"email"`
	OrganizationName string `json:"organisation_name"`
}
