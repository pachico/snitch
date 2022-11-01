package system

import (
	"os"
	"pachico/snitch/internal/domain" // replace with your actual project path
	"strings"
)

type EnvVarsReportRepository map[string]string

func (r *EnvVarsReportRepository) GetEnvVarsReport() (domain.EnvVarsReport, error) {
	report := domain.EnvVarsReport{} // assuming domain.EnvVarsReport is a compatible type

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		report[pair[0]] = pair[1]
	}

	return report, nil
}
