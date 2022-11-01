package handler

import (
	"net/http"
	"pachico/snitch/internal/domain/repository"

	"clevergo.tech/jsend"
	"github.com/labstack/echo/v4"
)

type FSReportHandler struct {
	Repository repository.FSReportRepositoryInterface
}

func (h FSReportHandler) Handle(c echo.Context) error {

	report, err := h.Repository.GetFSReport()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, jsend.NewFail(err.Error()))
	}

	return c.JSON(http.StatusOK, jsend.New(report))
}
