package domain

import "net/http"

// HTTPRequest holds the hostname used for an HTTP request.
type HTTPRequest struct {
	URL string `json:"url"` // Hostname to which the request is made
}

// HTTPResponse captures details about the response from an HTTP request.
type HTTPResponse struct {
	StatusCode int         `json:"status_code"` // HTTP status code of the response
	Size       int64       `json:"size"`        // Content length of the response body
	Header     http.Header `json:"headers"`     // HTTP headers returned with the response
	Body       string      `json:"body"`        // Response body as a string
}

// HTTPRequestReport groups together an HTTPRequest and its corresponding HTTPResponse.
type HTTPRequestReport struct {
	HTTPRequest  HTTPRequest  `json:"request"`
	HTTPResponse HTTPResponse `json:"response"`
}
