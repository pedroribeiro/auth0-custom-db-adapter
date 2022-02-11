package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pedroribeiro/users/config"
	"github.com/pedroribeiro/users/internal/app"
	"github.com/pedroribeiro/users/internal/driver/http"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {

	err := loadConfig(".env")

	if err != nil {
		exit(err)
	}
	// Create Http Server
	httpServer := http.New()

	// Starts App
	app.Start(httpServer.Router)

	httpServer.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	logrus.Info("received signal", <-c)
}

func loadConfig(filename string) error {
	err := config.ReadConfig(filename)
	if err != nil {
		return errors.Wrap(err, "read container")
	}
	return nil
}

func exit(err error) {
	logrus.Error(err)
	os.Exit(2)
}
