package routes_test

import (
	"testing"

	"pachico/snitch/internal/infrastructure/routes"

	"github.com/stretchr/testify/assert"
)

func TestRoutesExist(t *testing.T) {
	assert.Equal(t, "/", routes.RouteSystemReport)
	assert.Equal(t, "/envvars", routes.RouteEnvVarsReport)
	assert.Equal(t, "/fs", routes.RouteFSystemReport)
	assert.Equal(t, "/dnsresolution/:hostname", routes.RouteDNSResolutionReport)
	assert.Equal(t, "/httprequest/:hostname", routes.RouteHTTPRequestReport)
}
