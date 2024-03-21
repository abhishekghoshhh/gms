package model

import (
	"slices"
	"strings"
)

type IamProfileResponse struct {
	Id          string  `json:"id"`
	DisplayName string  `json:"displayName"`
	Groups      []Group `json:"groups"`
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

type Group struct {
	Display string `json:"display"`
}
