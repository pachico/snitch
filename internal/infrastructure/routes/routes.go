package routes

// Constants defining API endpoints for the system reports and functionalities.
const (
	RouteSystemReport        = "/"                        // Root route for system report
	RouteEnvVarsReport       = "/envvars"                 // Environment variables report
	RouteFSystemReport       = "/fs"                      // File System report
	RouteDNSResolutionReport = "/dnsresolution/:hostname" // DNS resolution report for a specific hostname
	RouteHTTPRequestReport   = "/httprequest/:hostname"   // HTTP request report for a specific hostname
)
