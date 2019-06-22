package server

import (
	"context"
	"net/http"

	"log"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func makeRouter(usePprof bool) *fasthttprouter.Router {
	router := &fasthttprouter.Router{
		RedirectTrailingSlash:  false,
		RedirectFixedPath:      false,
		HandleMethodNotAllowed: false,
		HandleOPTIONS:          true,

		PanicHandler: func(ctx *fasthttp.RequestCtx, i interface{}) {
			log.Printf("!PANIC! `%v`", i)
			ctx.SetStatusCode(http.StatusInternalServerError)
		},
	}

	// Add endpoints
	AddAgentRoutes(router)
	AddUtilityRoutes(router)

	if usePprof {
		// FIXME: pprof broke with new http library
		//router.HandleFunc("/debug/pprof/", pprof.Index)
		//router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		//router.HandleFunc("/debug/pprof/profile", pprof.Profile)
		//router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		//router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}

	return router
}

func StartHTTPServer(address string, logRequests bool, usePprof bool) *fasthttp.Server {
	// Wrap router in handler that captures panics
	router := makeRouter(usePprof)

	if logRequests {
		// FIXME: Logging broke with new http library
		//router = handlers.LoggingHandler(os.Stdout, router)
	}

	server := &fasthttp.Server{
		Handler: router.Handler,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Run server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(address); err != nil {
			log.Println(err)
		}
	}()

	return server
}

func StopHTTPServer(server *fasthttp.Server, existingConnectionTimeout time.Duration) {
	// Create a deadline to wait for existing connections to finish
	_, cancel := context.WithTimeout(context.Background(), existingConnectionTimeout)
	defer cancel()

	log.Println("Waiting for existing connections to finish...")
	if err := server.Shutdown(); err != nil {
		panic(err)
	}
}
