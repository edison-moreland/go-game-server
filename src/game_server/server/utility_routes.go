package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func AddUtilityRoutes(router *mux.Router) {
	// TODO Switch to using sub routers

	router.HandleFunc("/ping", pingHandler).Methods("GET")
}

// pingHandler is used by client to check that server is online
func pingHandler(w http.ResponseWriter, request *http.Request) {
	// TODO: Return server information. players online, etc
	defer request.Body.Close()

	w.WriteHeader(http.StatusNoContent)
	return
}
