package config

import (
	"os"
	"strconv"
	"testing"
)

func TestNewConfigDefaultPort(t *testing.T) {
	os.Unsetenv(envPortKey)

	config, err := New()
	if err != nil {
		t.Fatalf("Error should not have occurred: %v", err)
	}
	if config.GetPort() != defaultPort {
		t.Errorf("Expected port %d, got %d", defaultPort, config.GetPort())
	}
}

func TestNewConfigCustomPort(t *testing.T) {
	expectedPort := 8080
	os.Setenv(envPortKey, strconv.Itoa(expectedPort))
	defer os.Unsetenv(envPortKey)

	config, err := New()
	if err != nil {
		t.Fatalf("Error should not have occurred: %v", err)
	}
	if config.GetPort() != expectedPort {
		t.Errorf("Expected port %d, got %d", expectedPort, config.GetPort())
	}
}

func TestNewConfigInvalidPort(t *testing.T) {
	os.Setenv(envPortKey, "invalid")
	defer os.Unsetenv(envPortKey)

	_, err := New()
	if err == nil {
		t.Error("Expected an error due to invalid port, but got none")
	}
}
