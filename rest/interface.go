package rest

import (
	"context"
	"errors"
	"fmt"
	"github.com/bygui86/go-gorm/database"
	"time"

	"gopkg.in/logex.v1"
)

func NewRestInterface(dbInterface database.DbInterface) RestInterface {
	logex.Debug("Create new rest interface")

	cfg := loadConfig()

	server := &RestInterfaceImpl{
		config:      cfg,
		dbInterface: dbInterface,
	}

	server.setupRouter()
	server.setupHTTPServer()
	return server
}

func (r RestInterfaceImpl) Start() error {
	logex.Info("Start rest interface")

	if r.httpServer != nil && !r.running {
		go func() {
			err := r.httpServer.ListenAndServe()
			if err != nil {
				logex.Errorf("Error starting rest interface: %s", err.Error())
			}
		}()
		r.running = true
		logex.Infof("rest interface listening on port %d", r.config.RestPort)
		return nil
	}

	return errors.New("rest interface start failed: HTTP server not initialized or HTTP server already running")
}

func (r RestInterfaceImpl) Shutdown(timeout int) error {
	logex.Warn(fmt.Sprintf("Shutdown rest interface with timeout %d", timeout))

	if r.httpServer != nil && r.running {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		err := r.httpServer.Shutdown(ctx)
		if err != nil {
			logex.Errorf("Error shutting down rest interface: %s", err.Error())
		}
		r.running = false
		return nil
	}

	return errors.New("rest interface shutdown failed: HTTP server not initialized or HTTP server not running")
}
