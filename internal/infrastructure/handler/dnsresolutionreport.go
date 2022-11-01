package handler

import (
	"net/http"
	"pachico/snitch/internal/domain/repository"

	"clevergo.tech/jsend"
	"github.com/labstack/echo/v4"
)

type DNSResolutionHandler struct {
	Repository repository.DNSResolutionReportRepositoryInterface
}

func (h DNSResolutionHandler) Handle(c echo.Context) error {

	hostname := c.Param("hostname")

	report, err := h.Repository.GetDNSResolutionReport(hostname)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, jsend.NewFail(err.Error()))
	}

	return c.JSON(http.StatusOK, jsend.New(report))
}
