package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestConfString -
func TestConfString(t *testing.T) {
	testcases := []struct {
		name  string
		key   string
		value string
	}{
		{"Default Country ISO", "default.country_iso", "tw"},
		{"Default Time Zone", "default.time_zone", "+08:00"},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			value := ConfString(tc.key)
			assert.Equal(t, tc.value, value)
		})
	}
}

// TestConfInt -
func TestConfInt(t *testing.T) {
	testcases := []struct {
		name  string
		key   string
		value int
	}{
		{"App Port", "app_port", 8089},
		{"MySQL Max Connections", "mysql.max_connection", 2000},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			value := ConfInt(tc.key)
			assert.Equal(t, tc.value, value)
		})
	}
}

// TestConfBool -
func TestConfBool(t *testing.T) {
	testcases := []struct {
		name  string
		key   string
		value bool
	}{
		{"Logs stdout Only", "logs.stdout_only", true},
		{"Logs print log", "logs.print_log", true},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			value := ConfBool(tc.key)
			assert.Equal(t, tc.value, value)
		})
	}
}
