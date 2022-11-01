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

type MockWorkingEnvVarsReportRepository struct{}

func (r MockWorkingEnvVarsReportRepository) GetEnvVarsReport() (domain.EnvVarsReport, error) {
	report := domain.EnvVarsReport{
		"john": "lennon",
		"paul": "mccartney",
	}

	return report, nil
}

// If repository returns a report, handler should return it
func TestEnvVarsHandlerRespondsWithReportIfOK(t *testing.T) {

	// Arrange
	handler := handler.EnvVarsHandler{
		Repository: MockWorkingEnvVarsReportRepository{},
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
		`{"status":"success","data":{"john":"lennon","paul":"mccartney"}}
`,
		recorder.Body.String(),
	)
}

type MockFailingEnvVarsReportRepository struct{}

func (r MockFailingEnvVarsReportRepository) GetEnvVarsReport() (domain.EnvVarsReport, error) {
	return domain.EnvVarsReport{}, errors.New("My fake error")
}

func TestEnvVarsHandlerRespondsWithErrorIfRequired(t *testing.T) {

	// Arrange
	handler := handler.EnvVarsHandler{
		Repository: MockFailingEnvVarsReportRepository{},
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
