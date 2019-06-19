package server

import (
	"encoding/json"
	"game_server/db"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func AddAgentRoutes(router *mux.Router) {
	// TODO Switch to using sub routers

	// TODO: add deregister endpoint
	router.HandleFunc("/agents", getAgentsHandler).Methods("GET")
	router.HandleFunc("/agent", newAgentHandler).Methods("POST")
	router.HandleFunc("/agent/{id}", getAgentHandler).Methods("GET")
	router.HandleFunc("/agent/{id}", updateAgentHandler).Methods("PUT")
	router.HandleFunc("/agent/{id}", deleteAgentHandler).Methods("DELETE")
}

// getAgentsHandler returns the uuid of all agents registered to simulation
func getAgentsHandler(w http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	agents, err := db.GetAllAgents()
	if err != nil {
		log.Panic(err.Error())
	}

	if err := json.NewEncoder(w).Encode(agents); err != nil {
		log.Panic(err.Error())
	}
}

// newAgentHandler registers a new agent to the controller
func newAgentHandler(w http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var newAgent = db.Agent{}
	if request.ContentLength > 0 {
		if err := UnmarshalBody(request.Body, &newAgent); err != nil {
			log.Panic(err.Error())
		}
	}

	// Create new agent
	agent, err := db.NewAgent(newAgent)
	if err != nil {
		log.Panic(err.Error())
	}

	// Write JSON encoded model to response
	if err := json.NewEncoder(w).Encode(agent); err != nil {
		log.Panic(err.Error())
	}

}

// getAgentStateHandler returns the state of a specific agent
func getAgentHandler(w http.ResponseWriter, request *http.Request) {
	// Get id from url
	id, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get agent
	agent, err := db.GetAgent(id)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Write JSON encoded model to response
	if err := json.NewEncoder(w).Encode(agent); err != nil {
		log.Panic(err.Error())
	}
}

// updateAgentStateHandler returns the state of a specific agent
func updateAgentHandler(w http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	// Get id from url
	id, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// unmarshal request body
	var agentUpdate = db.Agent{}
	if err := UnmarshalBody(request.Body, &agentUpdate); err != nil {
		log.Panic(err.Error())
	}

	// Make sure ID is the same
	agentUpdate.ID = id

	// Update agent
	if err := db.UpdateAgent(agentUpdate); err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteAgentHandler(w http.ResponseWriter, request *http.Request) {
	// Get id from url
	id, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = db.DeleteAgent(id); err != nil {
		log.Panic(err)
	}
}
