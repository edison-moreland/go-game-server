package server

import (
	"context"
	"net/http/pprof"

	//"encoding/json"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func makeRouter(usePprof bool) *mux.Router {
	router := mux.NewRouter()

	// Add endpoints
	AddAgentRoutes(router)
	AddUtilityRoutes(router)

	if usePprof {
		router.HandleFunc("/debug/pprof/", pprof.Index)
		router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		router.HandleFunc("/debug/pprof/profile", pprof.Profile)
		router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}

	return router
}

func StartHTTPServer(address string, logRequests bool, usePprof bool) *http.Server {
	// Wrap router in handler that captures panics
	router := handlers.RecoveryHandler()(makeRouter(usePprof))
	//router := makeRouter()

	if logRequests {
		router = handlers.LoggingHandler(os.Stdout, router)
	}

	server := &http.Server{
		Addr:    address,
		Handler: router,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Run server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	return server
}

func StopHTTPServer(server *http.Server, existingConnectionTimeout time.Duration) {
	// Create a deadline to wait for existing connections to finish
	ctx, cancel := context.WithTimeout(context.Background(), existingConnectionTimeout)
	defer cancel()

	log.Println("Waiting for existing connections to finish...")
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
