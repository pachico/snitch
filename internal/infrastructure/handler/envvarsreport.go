package handler

import (
	"net/http"
	"pachico/snitch/internal/domain/repository"

	"clevergo.tech/jsend"
	"github.com/labstack/echo/v4"
)

type EnvVarsHandler struct {
	Repository repository.EnvVarsReportRepositoryInterface
}

func (h EnvVarsHandler) Handle(c echo.Context) error {

	report, err := h.Repository.GetEnvVarsReport()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, jsend.NewFail(err.Error()))
	}

	return c.JSON(http.StatusOK, jsend.New(report))
}
