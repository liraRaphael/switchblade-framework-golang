package util

import (
	"strconv"
)

func StringToBool(value string) bool {
	final, err := strconv.ParseBool(value)

	if err != nil {
		return false
	}

	return final
}
