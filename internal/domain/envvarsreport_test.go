package domain

import (
	"encoding/json"
	"testing"
)

func TestEnvVarsReportSerialization(t *testing.T) {
	report := EnvVarsReport{
		"PATH": "/usr/bin",
		"HOME": "/home/user",
	}

	bytes, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("Failed to serialize EnvVarsReport: %v", err)
	}

	jsonStr := string(bytes)
	expected := `{"HOME":"/home/user","PATH":"/usr/bin"}`
	if jsonStr != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, jsonStr)
	}
}
