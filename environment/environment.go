package environment

import (
	"os"
	"strings"
)

func GetEnvValue(value string) string {
	return os.Getenv(value)
}

func GetEnvValueOrDefault(value string, valueDefault string) string {
	final := GetEnvValue(value)

	if len(strings.TrimSpace(final)) == 0 {
		final = valueDefault
	}

	return final
}
