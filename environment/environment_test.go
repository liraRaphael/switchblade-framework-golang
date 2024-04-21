package environment

import (
	"testing"
)

func TestGetEnvValue(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvValue(tt.args.value); got != tt.want {
				t.Errorf("GetEnvValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvValueOrDefault(t *testing.T) {
	type args struct {
		value        string
		valueDefault string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvValueOrDefault(tt.args.value, tt.args.valueDefault); got != tt.want {
				t.Errorf("GetEnvValueOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
