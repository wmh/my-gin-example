package core

import (
	"testing"
)

// TestLog -
func TestLog(t *testing.T) {
	Log("test.log", H{"msg": "test log"})
}

// TestErrorLog -
func TestErrorLog(t *testing.T) {
	ErrorLog("test.error", H{"msg": "test errorlog", "code": 400})
}
