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

type MockWorkingDNSResolutionReportRepository struct{}

func (r MockWorkingDNSResolutionReportRepository) GetDNSResolutionReport(hostname string) (domain.DNSResolutionReport, error) {
	return domain.DNSResolutionReport{
		Hostname: "foobar",
		IPs: []domain.DNSResolution{
			{
				IP:              "1.2.3.4",
				IsPrivate:       true,
				IsGlobalUnicast: false,
				IsLoopback:      false,
				IsUnspecified:   false,
			}},
	}, nil
}

func TestDNSResolutionHandlerRespondsWithReportIfOK(t *testing.T) {

	// Arrange
	handler := handler.DNSResolutionHandler{
		Repository: MockWorkingDNSResolutionReportRepository{},
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
		`{"status":"success","data":{"hostname":"foobar","ips":[{"ip":"1.2.3.4","is_private":true,"is_global_unicast":false,"is_loopback":false,"is_unspecified":false}]}}
`,
		recorder.Body.String(),
	)
}

type MockFailingDNSResolutionReportRepository struct{}

func (r MockFailingDNSResolutionReportRepository) GetDNSResolutionReport(hostname string) (domain.DNSResolutionReport, error) {
	return domain.DNSResolutionReport{}, errors.New("My fake error")
}

func TestDNSResolutionHandlerRespondWithErrorIfRequired(t *testing.T) {

	// Arrange
	handler := handler.DNSResolutionHandler{
		Repository: MockFailingDNSResolutionReportRepository{},
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
