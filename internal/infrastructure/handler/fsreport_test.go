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

type MockWorkingFSReportRepository struct{}

func (r MockWorkingFSReportRepository) GetFSReport() (domain.FSReport, error) {
	return domain.FSReport{
		domain.File{
			Name:       "file1",
			Size:       1,
			IsDir:      false,
			Owner:      "root",
			Permission: "rwx",
		},
		domain.File{
			Name:       "file2",
			Size:       2,
			IsDir:      false,
			Owner:      "root",
			Permission: "rwx",
		},
	}, nil
}

func TestFSReportHandlerRespondsWithReportIfOK(t *testing.T) {
	// Arrange
	handler := handler.FSReportHandler{
		Repository: MockWorkingFSReportRepository{},
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
		`{"status":"success","data":[{"name":"file1","size":1,"is_dir":false,"owner":"root","permission":"rwx"},{"name":"file2","size":2,"is_dir":false,"owner":"root","permission":"rwx"}]}
`,
		recorder.Body.String(),
	)
}

type MockFailingFSReportRepository struct{}

func (r MockFailingFSReportRepository) GetFSReport() (domain.FSReport, error) {
	return domain.FSReport{}, errors.New("My fake error")
}

func TestFSReportHandlerRespondsWithErrorIfRequired(t *testing.T) {

	// Arrange
	handler := handler.FSReportHandler{
		Repository: MockFailingFSReportRepository{},
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
