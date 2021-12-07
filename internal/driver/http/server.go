package http

import (
	"context"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"github.com/sirupsen/logrus"
)

//nolint
var (
	nrapp newrelic.Application
)

//nolint
const (
	keyNrID          int           = iota
	TIMEOUT_SHUTDOWN time.Duration = 5 * time.Second
)

// Http Server
type HTTP struct {
	Router   *gin.Engine
	Listener *net.Listener
	Server   *http.Server
	Log      *logrus.Entry
}

// Listen for requests
func (c *HTTP) Run() {
	// Ready
	c.Log.Info("Listen on -> ", os.Getenv("HTTP_ADDR"))

	// Http server
	if err := c.Server.Serve(*c.Listener); err != nil && err != http.ErrServerClosed {
		c.Log.Fatal("Server closed unexpect")
	}
}

// Close http server
func (c *HTTP) Close() {
	// The context is used to inform the server it has 'timeout' to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_SHUTDOWN)
	defer cancel()

	// Shutdown server
	if err := c.Server.Shutdown(ctx); err != nil {
		c.Log.Error("Forced to shutdown: ", err)
	}
}

// NewServer creates an instance of Http Server
func New() *HTTP {
	// Looger Http
	log := logrus.WithFields(logrus.Fields{"module": "http"})

	// Create a listener tcp
	listener, err := net.Listen("tcp", os.Getenv("HTTP_ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	// Create router
	router := gin.Default()
	cfg := newrelic.NewConfig(os.Getenv("NEW_RELIC_APP_NAME"), os.Getenv("NEW_RELIC_TOKEN"))
	app, err := newrelic.NewApplication(cfg)
	if err != nil {
		log.Printf("failed to make new_relic app: %v", err)
	} else {
		router.Use(NewRelicMonitoring(app))
	}

	//nolint
	// Create server
	return &HTTP{
		Router:   router,
		Listener: &listener,
		Server:   &http.Server{Handler: router},
		Log:      log,
	}
}

const (
	// NewRelicTxnKey is the key used to retrieve the NewRelic Transaction from the context
	NewRelicTxnKey = "NewRelicTxnKey"
)

// NewRelicMonitoring is a middleware that starts a newrelic transaction, stores it in the context, then calls the next handler
func NewRelicMonitoring(app newrelic.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		txn := app.StartTransaction(ctx.Request.URL.Path, ctx.Writer, ctx.Request)
		defer txn.End()
		ctx.Set(NewRelicTxnKey, txn)
		ctx.Next()
	}
}
