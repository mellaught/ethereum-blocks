package app

import (
	"context"
	"net/http"
	"time"

	eth "github.com/mellaught/ethereum-blocks/src/ethereum"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type App struct {
	logger *logrus.Logger
	router *mux.Router
	Server *http.Server
	ethSRV *eth.EthereumSRV
}

// NewApp is initializes the app.
func NewApp(logger *logrus.Logger, explorerURL string) *App {
	// create new app
	a := &App{
		logger: logger,
		router: mux.NewRouter(),
		Server: &http.Server{},
		ethSRV: eth.CreateNewEthereumSRV(logger, explorerURL),
	}
	// set router && create ethereum service
	a.router = mux.NewRouter()
	a.setRouters()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	a.Server.Handler = handlers.CORS(headers, methods, origins)(a.router)

	return a
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("GET")
}

func (a *App) setRouters() {
	a.Get("/blocks/{input}", a.ethSRV.BlocksHandle)
}

// Run the app on it's router
func (a *App) Run(ctx context.Context, url string) {
	a.Server.Addr = url

	go func() {
		if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal(err)
		}
	}()
	a.logger.Infof("Ethereum block scanner service has started on %s\nPress ctrl + C to exit.", url)

	<-ctx.Done()

	a.logger.Infoln("Ethereum block scanner service has stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctxShutDown); err != nil {
		a.logger.Fatalf("Shutdown: %v\n", err)
	}
}
