package environment

import (
	"os"
	"strings"
)

// GetEnvValue returns the value of the environment variable 'value'.
//
// value: the name of the environment variable.
// string: the value of the environment variable.
func GetEnvValue(value string) string {
	return os.Getenv(value)
}

// GetEnvValueOrDefault returns the value of the environment variable 'value' if it is set,
// otherwise it returns the default value 'valueDefault'.
//
// value: the name of the environment variable.
// valueDefault: the default value to return if the environment variable is not set.
// string: the value of the environment variable, or the default value if the environment variable is not set.
func GetEnvValueOrDefault(value string, valueDefault string) string {
	final := GetEnvValue(value)

	if len(strings.TrimSpace(final)) == 0 {
		final = valueDefault
	}

	return final
}
