package mysql

import (
	"errors"
	"github.com/bygui86/go-gorm/utils"
)

const (
	hostEnv     = "MYSQL_HOST"
	portEnv     = "MYSQL_PORT"
	usernameEnv = "MYSQL_USERNAME"
	passwordEnv = "MYSQL_PASSWORD"
	paramsEnv   = "MYSQL_PARAMS"

	hostDefault     = ""
	portDefault     = 0
	usernameDefault = ""
	passwordDefault = ""
	paramsDefault   = ""
)

func loadConfig() (*internalConfig, error) {
	cfg := &internalConfig{
		host:     utils.GetStringEnv(hostEnv, hostDefault),
		port:     utils.GetIntEnv(portEnv, portDefault),
		username: utils.GetStringEnv(usernameEnv, usernameDefault),
		password: utils.GetStringEnv(passwordEnv, passwordDefault),
		params:   utils.GetStringEnv(paramsEnv, paramsDefault)}

	cfgErr := cfg.validateConfig()
	if cfgErr != nil {
		return nil, cfgErr
	}

	return cfg, nil
}

func (c *internalConfig) validateConfig() error {
	if c.host == hostDefault {
		return errors.New("host must be specified")
	}
	if c.port == portDefault {
		return errors.New("port must be specified")
	}
	if c.username == usernameDefault {
		return errors.New("username must be specified")
	}
	if c.password == passwordDefault {
		return errors.New("password must be specified")
	}
	return nil
}
