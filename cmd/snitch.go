package main

import (
	"fmt"
	"os"
	"pachico/snitch/internal/config"
	"pachico/snitch/internal/infrastructure/handler"
	"pachico/snitch/internal/infrastructure/repository/system"
	"pachico/snitch/internal/infrastructure/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.Lvl(2)) // log.INFO
	e.Use(middleware.Logger())

	conf, err := config.New()
	if err != nil {
		e.Logger.Errorf("Failed to load configuration: %v", err)
		os.Exit(1)
	}

	envVarsHandler := handler.EnvVarsHandler{
		Repository: &system.EnvVarsReportRepository{},
	}

	systemReportHandler := handler.SystemReportHandler{
		Repository: &system.SystemReportRepository{},
	}

	fileReportHandler := handler.FSReportHandler{
		Repository: &system.FSReportRepository{},
	}

	dnsResolutionReportHandler := handler.DNSResolutionHandler{
		Repository: &system.DNSResolutionReportRepository{},
	}

	httpRequestReportHandler := handler.HTTPRequestReportHandler{
		Repository: &system.HTTPRequestReportRepository{},
	}

	e.GET(routes.RouteSystemReport, systemReportHandler.Handle)
	e.GET(routes.RouteEnvVarsReport, envVarsHandler.Handle)
	e.GET(routes.RouteFSystemReport, fileReportHandler.Handle)
	e.GET(routes.RouteDNSResolutionReport, dnsResolutionReportHandler.Handle)
	e.GET(routes.RouteHTTPRequestReport, httpRequestReportHandler.Handle)

	port := fmt.Sprintf(":%d", conf.GetPort())
	e.Logger.Fatal(e.Start(port))
}
