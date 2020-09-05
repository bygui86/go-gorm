package monitoring

import (
	"net/http"

	"github.com/gorilla/mux"
)

type MonitorInterface interface {
	Start()
	Shutdown(timeout int)
}

type MonitorInterfaceImpl struct {
	config     *config
	router     *mux.Router
	httpServer *http.Server
	running    bool
}

type config struct {
	restHost string
	restPort int
}
