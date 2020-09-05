package monitoring

import (
	"gopkg.in/logex.v1"

	"github.com/bygui86/go-gorm/utils"
)

const (
	monitorHostEnvVar = "MONITOR_REST_HOST"
	monitorPortEnvVar = "MONITOR_REST_PORT"

	monitorHostDefault = "localhost"
	monitorPortDefault = 9090
)

func loadConfig() *config {
	logex.Debug("Load monitoring configurations")
	return &config{
		restHost: utils.GetStringEnv(monitorHostEnvVar, monitorHostDefault),
		restPort: utils.GetIntEnv(monitorPortEnvVar, monitorPortDefault),
	}
}
