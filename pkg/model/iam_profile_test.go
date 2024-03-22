package model

import (
	"testing"
)

func TestIamProfileResponse(t *testing.T) {
	t.Run("should return matching groups from the requestedGroups", func(t *testing.T) {
		groups := []Group{
			{"firstGroup"},
			{"secondGroup"},
		}
		iamProfileResponse := IamProfileResponse{
			Id:          "12",
			DisplayName: "this is a display name",
			Groups:      groups,
		}

		matchingProfiles := iamProfileResponse.GetMatchingGroups([]string{"firstGroup", "test"})
		wanted := "firstGroup\n"

		if matchingProfiles != wanted {
			t.Error("Result was incorrect got:", matchingProfiles, "want:", wanted)
		}
	})

	t.Run("should return empty string if no iam profile groups are available", func(t *testing.T) {
		iamProfileResponse := IamProfileResponse{
			Id:          "12",
			DisplayName: "this is a display name",
			Groups:      make([]Group, 0),
		}

		matchingProfiles := iamProfileResponse.GetMatchingGroups([]string{"firstGroup"})
		wanted := ""

		if matchingProfiles != wanted {
			t.Error("Result was incorrect got:", matchingProfiles, "want:", wanted)
		}
	})

	t.Run("should return available iam profile groups if no requestedGroups are available", func(t *testing.T) {
		groups := []Group{
			{"firstGroup"},
			{"secondGroup"},
		}

		iamProfileResponse := IamProfileResponse{
			Id:          "12",
			DisplayName: "this is a display name",
			Groups:      groups,
		}

		matchingProfiles := iamProfileResponse.GetMatchingGroups(make([]string, 0))
		wanted := "firstGroup\nsecondGroup\n"

		if matchingProfiles != wanted {
			t.Error("Result was incorrect got:", matchingProfiles, "want:", wanted)
		}
	})
}
