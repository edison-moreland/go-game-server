package server

import (
	"game_server/agent_store"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
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
	agents := agent_store.GetAllAgents()

	_, _, err := easyjson.MarshalToHTTPResponseWriter(agents, w)
	if err != nil {
		log.Panic(err.Error())
	}
}

// newAgentHandler registers a new agent to the controller
func newAgentHandler(w http.ResponseWriter, request *http.Request) {
	var newAgent = agent_store.Agent{}
	if request.ContentLength > 0 {
		if err := easyjson.UnmarshalFromReader(request.Body, &newAgent); err != nil {
			log.Panic(err.Error())
		}
	}

	// Create new agent
	agent := agent_store.NewAgent(newAgent)

	// Write JSON encoded model to response
	_, _, err := easyjson.MarshalToHTTPResponseWriter(agent, w)
	if err != nil {
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
	agent, err := agent_store.GetAgent(id)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Write JSON encoded model to response
	_, _, err = easyjson.MarshalToHTTPResponseWriter(agent, w)
	if err != nil {
		log.Panic(err.Error())
	}
}

// updateAgentStateHandler returns the state of a specific agent
func updateAgentHandler(w http.ResponseWriter, request *http.Request) {
	// Get id from url
	id, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// unmarshal request body
	var agentUpdate = agent_store.Agent{}
	if err := easyjson.UnmarshalFromReader(request.Body, &agentUpdate); err != nil {
		log.Panic(err.Error())
	}

	// Make sure ID is the same
	agentUpdate.ID = id

	// Update agent
	agent_store.UpdateAgent(agentUpdate)

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

	agent_store.DeleteAgent(id)
}
