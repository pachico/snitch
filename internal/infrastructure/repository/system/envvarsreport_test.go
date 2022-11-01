package system_test

import (
	"os"
	"pachico/snitch/internal/infrastructure/repository/system"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvVarsReportRepositoryReturnsProperStructIfEmpty(t *testing.T) {
	// Arrange
	repo := system.EnvVarsReportRepository{}
	envVarKey := "asdya7sdynaisdai8sucnliausdoaysndclia"
	defer os.Unsetenv(envVarKey)
	os.Setenv(envVarKey, "foobar")

	// Act
	report, err := repo.GetEnvVarsReport()
	// Assert
	assert.Nil(t, err)
	assert.NotEmpty(t, report)

	envVarFound := false
	for k, v := range report {
		if k == envVarKey {
			assert.Equal(t, "foobar", v)
			envVarFound = true
		}
	}

	assert.True(t, envVarFound)
}
