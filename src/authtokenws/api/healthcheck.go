package response

type HealthCheckResult struct {
	Healthy        bool   `json:"healthy"`
	Message        string `json:"message,omitempty"`
}

type HealthCheck struct {
	CheckType      HealthCheckResult `json:"mysql"`
}

