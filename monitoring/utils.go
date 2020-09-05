package monitoring

import (
	"fmt"
	"gopkg.in/logex.v1"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// endpoints
	rootEndpoint = "/metrics"

	// server parameters
	httpServerHostFormat          = "%s:%d"
	httpServerWriteTimeoutDefault = time.Second * 15
	httpServerReadTimeoutDefault  = time.Second * 15
	httpServerIdleTimeoutDefault  = time.Second * 60
)

func (m *MonitorInterfaceImpl) setupRouter() {
	logex.Debugf("Setup new monitoring router")

	m.router = mux.NewRouter().StrictSlash(true)

	m.router.Handle(rootEndpoint, promhttp.Handler())
}

func (m *MonitorInterfaceImpl) newHTTPServer() {
	logex.Debugf("Setup new monitoring HTTP server on port %d...", m.config.restPort)

	if m.config != nil {
		m.httpServer = &http.Server{
			Addr:    fmt.Sprintf(httpServerHostFormat, m.config.restHost, m.config.restPort),
			Handler: m.router,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: httpServerWriteTimeoutDefault,
			ReadTimeout:  httpServerReadTimeoutDefault,
			IdleTimeout:  httpServerIdleTimeoutDefault,
		}
		return
	}

	logex.Error("Monitoring HTTP server creation failed: monitoring configurations not loaded")
}
