package rest

import (
	"github.com/bygui86/go-rest/utils"
	"gopkg.in/logex.v1"
)

const (
	restHostEnvVar = "REST_HOST"
	restPortEnvVar = "REST_PORT"

	restHostEnvVarDefault = "localhost"
	restPortEnvVarDefault = 8080
)

func loadConfig() *config {
	logex.Debug("Load configurations")
	return &config{
		RestHost: utils.GetStringEnv(restHostEnvVar, restHostEnvVarDefault),
		RestPort: utils.GetIntEnv(restPortEnvVar, restPortEnvVarDefault),
	}
}
