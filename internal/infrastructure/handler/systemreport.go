package handler

import (
	"net/http"
	"pachico/snitch/internal/domain/repository"

	"clevergo.tech/jsend"
	"github.com/labstack/echo/v4"
)

type SystemReportHandler struct {
	Repository repository.SystemReportRepositoryInterface
}

func (h SystemReportHandler) Handle(c echo.Context) error {

	report, err := h.Repository.GetSystemReport()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, jsend.New(report))
}
