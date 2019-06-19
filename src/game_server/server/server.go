package server

import (
	"context"
	"encoding/json"
	"github.com/gorilla/handlers"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func contentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func makeRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(contentTypeMiddleware)

	// Add endpoints
	AddAgentRoutes(router)
	AddUtilityRoutes(router)

	return router
}

func StartHTTPServer(address string, logRequests bool) *http.Server {
	// Wrap router in handler that captures panics
	router := handlers.RecoveryHandler()(makeRouter())
	//router := makeRouter()

	if logRequests {
		router = handlers.LoggingHandler(os.Stdout, makeRouter())
	}

	server := &http.Server{
		Addr:    address,
		Handler: router,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
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

func UnmarshalBody(body io.Reader, bodyObject interface{}) error {
	// Read request body
	rawBody, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	// Unmarshal json from rawBody
	if err = json.Unmarshal(rawBody, &bodyObject); err != nil {
		return err
	}

	return nil
}
