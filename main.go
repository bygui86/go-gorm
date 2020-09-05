package main

import (
	"fmt"
	"github.com/bygui86/go-gorm/database"
	"github.com/bygui86/go-gorm/monitoring"
	"github.com/bygui86/go-gorm/rest"
	"gopkg.in/logex.v1"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	restInt    rest.RestInterface
	monitorInt monitoring.MonitorInterface
)

func main() {
	logex.Info("Start go-gorm")

	dbInterface := startDbConnection()

	restInt = startRestServer(dbInterface)

	//monitorInt = startMonitoringServer()

	startSysCallChannel()

	shutdownAndWait(3)
}

func startDbConnection() database.DbInterface {
	logex.Info("Create new database connection")

	db, newErr := database.NewDbInterface()
	if newErr != nil {
		logex.Fatal(newErr)
	}
	logex.Debug("Database interface successfully created")

	openErr := db.OpenConnection()
	if openErr != nil {
		logex.Fatal(openErr)
	}
	logex.Debug("Database connection successfully opened")

	initErr := db.InitSchema()
	if initErr != nil {
		logex.Fatal(initErr)
	}
	logex.Debug("Database schema successfully initialized")

	return db
}

func startRestServer(dbInterface database.DbInterface) rest.RestInterface {
	logex.Info("Start rest server")

	server := rest.NewRestInterface(dbInterface)
	logex.Debug("rest server successfully created")

	err := server.Start()
	if err != nil {
		logex.Fatal(err)
	}
	logex.Debug("rest server successfully started")

	return server
}

func startMonitoringServer() monitoring.MonitorInterface {
	logex.Info("Start monitoring")
	server := monitoring.NewMonitorInterface()
	logex.Debug("Monitoring server successfully created")

	server.Start()
	logex.Debug("Monitoring successfully started")

	return server
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(timeout int) {
	logex.Warn(fmt.Sprintf("Termination signal received, timeout %d", timeout))

	if restInt != nil {
		err := restInt.Shutdown(timeout)
		if err != nil {
			logex.Errorf("Error during rest interface shutdown: %s", err.Error())
		}
	}

	if monitorInt != nil {
		monitorInt.Shutdown(timeout)
	}

	time.Sleep(time.Duration(timeout+1) * time.Second)
}
