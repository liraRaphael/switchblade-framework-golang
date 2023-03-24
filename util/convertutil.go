package util

import (
	"strconv"
)

func StringToBool(value string) bool {
	final, err := strconv.ParseBool(GetEnvValueOrDefault("application.healthcheck.active", "true"))

	if err != nil {
		return false
	}

	return final
}
