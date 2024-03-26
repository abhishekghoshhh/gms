package model

type IamProfileListResponse struct {
	TotalResults int                  `json:"totalResults"`
	ItemsPerPage int                  `json:"itemsPerPage"`
	StartIndex   int                  `json:"startIndex"`
	Resources    []IamProfileResponse `json:"Resources"`
}
