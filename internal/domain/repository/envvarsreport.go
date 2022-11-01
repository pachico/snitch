package repository

import "pachico/snitch/internal/domain"

type EnvVarsReportRepositoryInterface interface {
	GetEnvVarsReport() (domain.EnvVarsReport, error)
}
