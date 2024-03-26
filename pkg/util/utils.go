package util

import "strconv"

func Bool(val string) bool {
	if b, err := strconv.ParseBool(val); err != nil {
		return false
	} else {
		return b
	}
}
