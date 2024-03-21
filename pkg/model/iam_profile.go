package model

type IamProfileResponse struct {
	Id          string  `json:"id"`
	DisplayName string  `json:"displayName"`
	Groups      []Group `json:"groups"`
}

type Group struct {
	Display string `json:"display"`
}
