package domain

import (
	"encoding/json"
	"testing"
)

func TestFSReportSerialization(t *testing.T) {
	fsReport := FSReport{
		{Name: "testfile.txt", Size: 1024, IsDir: false},
		{Name: "data", Size: 2048, IsDir: true},
	}

	bytes, err := json.Marshal(fsReport)
	if err != nil {
		t.Fatalf("Failed to serialize FilesReport: %v", err)
	}

	jsonStr := string(bytes)
	expected := `[{"name":"testfile.txt","size":1024,"is_dir":false,"owner":"","permission":""},{"name":"data","size":2048,"is_dir":true,"owner":"","permission":""}]`
	if jsonStr != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, jsonStr)
	}
}
