package model

import (
	"slices"
	"strings"
)

type IamProfileResponse struct {
	Id               string           `json:"id"`
	DisplayName      string           `json:"displayName"`
	Groups           []Group          `json:"groups"`
	IndigoUserSchema IndigoUserSchema `json:"urn:indigo-dc:scim:schemas:IndigoUser"`
}

func (profile *IamProfileResponse) GetMatchingGroups(requestedGroups []string) string {
	groupsList := profile.Groups
	if len(groupsList) == 0 {
		return ""
	} else if len(requestedGroups) == 0 {
		var groups strings.Builder
		for _, group := range groupsList {
			groups.WriteString(group.Display)
			groups.WriteString("\n")
		}
		return groups.String()
	} else {
		var groups strings.Builder
		filteredGroups := profile.getUniqueGroups(requestedGroups)
		for _, group := range groupsList {
			if slices.Contains(filteredGroups, group.Display) {
				groups.WriteString(group.Display)
				groups.WriteString("\n")
			}
		}
		return groups.String()
	}
}

func (*IamProfileResponse) getUniqueGroups(requestedGroups []string) []string {
	filteredGroups := make([]string, 0)
	for _, requestedGroupName := range requestedGroups {
		if !slices.Contains(filteredGroups, requestedGroupName) {
			filteredGroups = append(filteredGroups, requestedGroupName)
		}
	}
	return filteredGroups
}

func (profile *IamProfileResponse) HasMatchingCert(subjectDn, clientCert string) bool {
	if len(profile.IndigoUserSchema.Certificates) == 0 {
		return false
	}
	for _, cert := range profile.IndigoUserSchema.Certificates {
		if cert.SubjectDn == subjectDn && cert.PemEncodedCertificate == clientCert {
			return true
		}
	}
	return false
}

type Group struct {
	Display string `json:"display"`
}

type IndigoUserSchema struct {
	Certificates []UserCertificate `json:"certificates"`
}
type UserCertificate struct {
	Primary               bool   `json:"primary"`
	SubjectDn             string `json:"subjectDn"`
	IssuerDn              string `json:"issuerDn"`
	PemEncodedCertificate string `json:"pemEncodedCertificate"`
	Display               string `json:"display"`
	Created               string `json:"created"`
	LastModified          string `json:"lastModified"`
	HasProxyCertificate   bool   `json:"hasProxyCertificate"`
}
