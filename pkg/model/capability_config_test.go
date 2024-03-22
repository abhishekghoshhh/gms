package model

import (
	"testing"
)

func TestCapabilitiesConfig(t *testing.T) {
	t.Run("should return the value for given key", func(t *testing.T) {
		config := NewCapabilitiesConfig(Entry{key: "group", value: "new"})

		value, _ := config.Get("group")
		wanted := "new"

		if value != wanted {
			t.Error("Result was incorrect, got:", value, "want:", wanted)
		}
	})

	t.Run("should return false if value for given key is not available", func(t *testing.T) {
		config := NewCapabilitiesConfig(Entry{key: "group", value: "new"})

		_, ok := config.Get("groups")
		if ok != false {
			t.Error("Result was incorrect, got:", ok, "want:", false)
		}
	})
}
