package util

import "testing"

func TestBool(t *testing.T) {
	t.Run("should get true", func(t *testing.T) {
		val := Bool("true")

		if val != true {
			t.Error("Result was incorrect got:", val, "want:", true)
		}
	})

	t.Run("should get false", func(t *testing.T) {
		val := Bool("false")

		if val != false {
			t.Error("Result was incorrect got:", val, "want:", false)
		}
	})

	t.Run("should parse valid boolean", func(t *testing.T) {
		value := Bool("True")

		if value != true {
			t.Error("Result was incorrect got:", value, "want:", true)
		}
	})

	t.Run("should return false for invalid bool string", func(t *testing.T) {
		value := Bool("TrUe")

		if value != false {
			t.Error("Result was incorrect got:", value, "want:", false)
		}
	})
}
