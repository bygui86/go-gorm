package monitoring

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/logex.v1"
)

func NewMonitorInterface() MonitorInterface {
	logex.Info("Create new monitoring server")

	cfg := loadConfig()
	server := &MonitorInterfaceImpl{
		config: cfg,
	}
	server.setupRouter()
	server.newHTTPServer()
	return server
}

func (m *MonitorInterfaceImpl) Start() {
	logex.Info("Start monitoring server")

	if m.httpServer != nil && !m.running {
		go func() {
			err := m.httpServer.ListenAndServe()
			if err != nil {
				logex.Errorf("Monitoring server start failed: %s", err.Error())
			}
		}()
		m.running = true
		logex.Infof("Monitoring server listen on port %d", m.config.restPort)
		return
	}

	logex.Error("Monitoring server start failed: HTTP server not initialized or HTTP server already running")
}

func (m *MonitorInterfaceImpl) Shutdown(timeout int) {
	logex.Warn(fmt.Sprintf("Shutdown monitoring server, timeout %d", timeout))

	if m.httpServer != nil && m.running {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		err := m.httpServer.Shutdown(ctx)
		if err != nil {
			logex.Errorf("Monitoring server shutdown failed: %s", err.Error())
		}
		m.running = false
		return
	}

	logex.Error("Monitoring server shutdown failed: HTTP server not initialized or HTTP server not running")
}
