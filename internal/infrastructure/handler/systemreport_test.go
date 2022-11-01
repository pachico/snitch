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

type MockFailingSystemReportRepository struct{}

func (r MockFailingSystemReportRepository) GetSystemReport() (domain.SystemReport, error) {
	return domain.SystemReport{}, errors.New("My fake error")
}

func TestSystemReportHandlerRespondsWithErrorIfRequired(t *testing.T) {

	// Arrange
	handler := handler.SystemReportHandler{
		Repository: MockFailingSystemReportRepository{},
	}
	echo := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/", nil)
	recorder := httptest.NewRecorder()
	context := echo.NewContext(request, recorder)

	// Act
	handler.Handle(context)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "\"My fake error\"\n", recorder.Body.String())
}

type MockWorkingSystemReportRepository struct{}

func (r MockWorkingSystemReportRepository) GetSystemReport() (domain.SystemReport, error) {
	return domain.SystemReport{
		Meta:    "meta",
		Node:    "node",
		OS:      "os",
		Kernel:  "kernel",
		Product: "product",
		Board:   "board",
		Chassis: "chassis",
		BIOS:    "bios",
		CPU:     "cpu",
		Memory:  "memory",
		Storage: []interface{}{"storage1", "storage2"},
		Network: []interface{}{"network1", "network2"},
	}, nil
}
func TestSystemReportHandlerRespondsWithReportIfOK(t *testing.T) {

	// Arrange
	handler := handler.SystemReportHandler{
		Repository: MockWorkingSystemReportRepository{},
	}
	echo := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/", nil)
	recorder := httptest.NewRecorder()
	context := echo.NewContext(request, recorder)

	// Act
	handler.Handle(context)

	// Assert
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t,
		`{"status":"success","data":{"sysinfo":"meta","node":"node","os":"os","kernel":"kernel","product":"product","board":"board","chassis":"chassis","bios":"bios","cpu":"cpu","memory":"memory","storage":["storage1","storage2"],"network":["network1","network2"]}}
`,
		recorder.Body.String(),
	)
}
