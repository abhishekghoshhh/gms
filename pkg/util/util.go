package util

import (
	"strconv"
)

func Bool(val string) bool {
	if b, err := strconv.ParseBool(val); err != nil {
		return false
	} else {
		return b
	}
}

func Int(val string) int {
	if num, err := strconv.Atoi(val); err != nil {
		return 0
	} else {
		return num
	}
}
