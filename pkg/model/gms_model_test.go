package model

import "testing"

func TestGmsModel(t *testing.T) {
	t.Run("should return true if token is available", func(t *testing.T) {
		model := GmsModel{token: "token", groups: nil, subjectDn: "subject", clientCert: "client_cert"}
		hasToken := model.HasToken()

		if hasToken != true {
			t.Error("Result was incorrect got:", hasToken, "want:", true)
		}
	})

	t.Run("should retun true if model has certificates", func(t *testing.T) {
		model := GmsModel{token: "token", groups: nil, subjectDn: "subject", clientCert: "client_cert"}
		hasToken := model.HasToken()

		if hasToken != true {
			t.Error("Result was incorrect got:", hasToken, "want:", true)
		}
	})

	t.Run("should return true if model has groups", func(t *testing.T) {
		model := GmsModel{
			token:      "token",
			groups:     []string{"first"},
			subjectDn:  "subject",
			clientCert: "client_cert",
		}
		hasToken := model.HasToken()

		if hasToken != true {
			t.Error("Result was incorrect got:", hasToken, "want:", true)
		}
	})

	t.Run("should return available groups if model has groups", func(t *testing.T) {
		model := GmsModel{
			token:      "token",
			groups:     []string{"first"},
			subjectDn:  "subject",
			clientCert: "client_cert",
		}
		groups := model.Groups()

		if groups[0] != "first" {
			t.Error("Result was incorrect got:", groups, "want:", []string{"first"})
		}
	})

	t.Run("should return token", func(t *testing.T) {
		model := GmsModel{token: "token", groups: nil, subjectDn: "subject", clientCert: "client_cert"}
		token := model.Token()
		wanted := "token"

		if token != wanted {
			t.Error("Result was incorrect got:", token, "want:", wanted)
		}
	})
}
