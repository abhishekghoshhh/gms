package model_test

import (
	"github.com/abhishekghoshhh/gms/pkg/model"
	"testing"
)

func TestPasswordGrantFlowConfigTest(t *testing.T) {
	t.Run("should create password grant flow config with given data", func(t *testing.T) {
		passwordGrantConfig := model.NewPasswordGrantFlowConfig("True", "username", "password", "id122", "clientSecret")

		if passwordGrantConfig.IsActive() != true {
			t.Error("Result was incorrect got:", passwordGrantConfig.IsActive(), "want:", true)
		}
	})

	t.Run("should set isActive to false if config is invalid", func(t *testing.T) {
		passwordGrantConfig := model.NewPasswordGrantFlowConfig("TRRue", "username", "password", "id122", "clientSecret")
		if passwordGrantConfig.IsActive() != false {
			t.Error("Result was incorrect got:", passwordGrantConfig.IsActive(), "want:", false)
		}
	})

	t.Run("should parse valid bool", func(t *testing.T) {
		value := model.Bool("True")

		if value != true {
			t.Error("Result was incorrect got:", value, "want:", true)
		}
	})

	t.Run("should return false for invalid bool string", func(t *testing.T) {
		value := model.Bool("FASLe")

		if value != false {
			t.Error("Result was incorrect got:", value, "want:", false)
		}
	})
}
