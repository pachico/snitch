package domain

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestHTTPRequestReportSerialization(t *testing.T) {
	request := HTTPRequest{
		URL: "example.com",
	}

	responseHeaders := http.Header{}
	responseHeaders.Add("Content-Type", "application/json")

	response := HTTPResponse{
		StatusCode: 200,
		Size:       1234,
		Header:     responseHeaders,
		Body:       `{"message": "success"}`,
	}

	report := HTTPRequestReport{
		HTTPRequest:  request,
		HTTPResponse: response,
	}

	bytes, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("Failed to serialize HTTPRequestReport: %v", err)
	}

	jsonStr := string(bytes)
	expected := `{"request":{"url":"example.com"},"response":{"status_code":200,"size":1234,"headers":{"Content-Type":["application/json"]},"body":"{\"message\": \"success\"}"}}`
	if jsonStr != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, jsonStr)
	}
}
