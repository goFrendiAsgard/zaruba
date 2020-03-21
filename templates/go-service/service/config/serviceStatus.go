package config

import "net/http"

// ServiceStatus represent liveness and readiness of service
type ServiceStatus struct {
	IsAlive bool
	IsReady bool
}

// GetLivenessHTTPCode get HTTP Code represent liveness status
func (s *ServiceStatus) GetLivenessHTTPCode() (httpCode int) {
	if s.IsAlive {
		return http.StatusOK
	}
	return http.StatusInternalServerError
}

// GetReadinessHTTPCode get HTTP Code represent readiness status
func (s *ServiceStatus) GetReadinessHTTPCode() (httpCode int) {
	if s.IsReady {
		return http.StatusOK
	}
	return http.StatusInternalServerError
}
