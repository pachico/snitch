package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"pachico/snitch/internal/domain"
	"pachico/snitch/internal/infrastructure/handler"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockWorkingHTTPRequestReportRepository struct{}

func (r MockWorkingHTTPRequestReportRepository) GetHTTPRequestReport(hostname string) (domain.HTTPRequestReport, error) {
	return domain.HTTPRequestReport{
		HTTPRequest: domain.HTTPRequest{
			URL: "localhost",
		},
		HTTPResponse: domain.HTTPResponse{
			StatusCode: 200,
			Size:       100,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: "Hello world",
		},
	}, nil
}

func TestHTTPRequestHandlerRespondsWithReportIfOK(t *testing.T) {

	// Arrange
	handler := handler.HTTPRequestReportHandler{
		Repository: MockWorkingHTTPRequestReportRepository{},
	}
	echo := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/", nil)
	recorder := httptest.NewRecorder()
	context := echo.NewContext(request, recorder)

	// Act
	handler.Handle(context)

	// Assert
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(
		t,
		`{"status":"success","data":{"request":{"url":"localhost"},"response":{"status_code":200,"size":100,"headers":{"Content-Type":["application/json"]},"body":"Hello world"}}}
`,
		recorder.Body.String(),
	)
}

type MockFailingHTTPRequestReportRepository struct{}

func (r MockFailingHTTPRequestReportRepository) GetHTTPRequestReport(hostname string) (domain.HTTPRequestReport, error) {
	return domain.HTTPRequestReport{}, errors.New("My fake error")
}

func TestHTTPRequestHandlerRespondsWithErrorIfRequired(t *testing.T) {

	// Arrange
	handler := handler.HTTPRequestReportHandler{
		Repository: MockFailingHTTPRequestReportRepository{},
	}
	echo := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/", nil)
	recorder := httptest.NewRecorder()
	context := echo.NewContext(request, recorder)

	// Act
	handler.Handle(context)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(
		t,
		`{"status":"fail","data":"My fake error"}
`,
		recorder.Body.String(),
	)
}
