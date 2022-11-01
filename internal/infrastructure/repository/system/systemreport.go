package system

import (
	"pachico/snitch/internal/domain"

	"github.com/zcalusic/sysinfo"
)

type SystemReportRepository struct {
}

func (sir *SystemReportRepository) GetSystemReport() (domain.SystemReport, error) {

	var si sysinfo.SysInfo

	si.GetSysInfo()

	info := domain.NewSystemReport(si)

	return *info, nil
}
