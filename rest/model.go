package rest

import (
	"github.com/bygui86/go-gorm/database"
	"net/http"

	"github.com/gorilla/mux"
)

type RestInterface interface {
	Start() error
	Shutdown(timeout int) error
}

type RestInterfaceImpl struct {
	config      *config
	router      *mux.Router
	httpServer  *http.Server
	dbInterface database.DbInterface
	running     bool
}

type config struct {
	RestHost string
	RestPort int
}

type errorResponse struct {
	Request string `json:"request"`
	Message string `json:"message"`
}
