package handler

import (
	"net/http"
	"pachico/snitch/internal/domain/repository"

	"clevergo.tech/jsend"
	"github.com/labstack/echo/v4"
)

type HTTPRequestReportHandler struct {
	Repository repository.HTTPRequestReportRepositoryInterface
}

func (h HTTPRequestReportHandler) Handle(c echo.Context) error {

	hostname := c.Param("hostname")

	report, err := h.Repository.GetHTTPRequestReport(hostname)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, jsend.NewFail(err.Error()))
	}

	return c.JSON(http.StatusOK, jsend.New(report))
}
