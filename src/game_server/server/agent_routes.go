package server

import (
	"game_server/agent_store"
	"github.com/buaazp/fasthttprouter"
	"github.com/google/uuid"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func AddAgentRoutes(router *fasthttprouter.Router) {
	// TODO Switch to using sub routers

	// TODO: add deregister endpoint
	router.GET("/agents", getAgentsHandler)
	router.POST("/agent", newAgentHandler)
	router.GET("/agent/:id", getAgentHandler)
	router.PUT("/agent/:id", updateAgentHandler)
	router.DELETE("/agent/:id", deleteAgentHandler)
}

// getAgentsHandler returns the uuid of all agents registered to simulation
func getAgentsHandler(ctx *fasthttp.RequestCtx) {
	agents := agent_store.GetAllAgents()

	_, err := easyjson.MarshalToWriter(agents, ctx)
	if err != nil {
		log.Panic(err.Error())
	}
}

// newAgentHandler registers a new agent to the controller
func newAgentHandler(ctx *fasthttp.RequestCtx) {
	var newAgent = agent_store.Agent{}
	if ctx.Request.Header.ContentLength() > 0 {
		if err := easyjson.Unmarshal(ctx.PostBody(), &newAgent); err != nil {
			log.Panic(err.Error())
		}
	}

	// Create new agent
	agent := agent_store.NewAgent(newAgent)

	// Write JSON encoded model to response
	_, err := easyjson.MarshalToWriter(agent, ctx)
	if err != nil {
		log.Panic(err.Error())
	}

}

// getAgentStateHandler returns the state of a specific agent
func getAgentHandler(ctx *fasthttp.RequestCtx) {
	// Get id from url
	id, err := uuid.Parse(ctx.UserValue("id").(string))
	if err != nil {
		log.Print(err.Error())
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	// Get agent
	agent, err := agent_store.GetAgent(id)
	if err != nil {
		log.Print(err.Error())
		ctx.SetStatusCode(http.StatusNotFound)
		return
	}

	// Write JSON encoded model to response
	_, err = easyjson.MarshalToWriter(agent, ctx)
	if err != nil {
		log.Panic(err.Error())
	}
}

// updateAgentStateHandler returns the state of a specific agent
func updateAgentHandler(ctx *fasthttp.RequestCtx) {
	// Get id from url
	id, err := uuid.Parse(ctx.UserValue("id").(string))
	if err != nil {
		log.Print(err.Error())
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	// unmarshal request body
	var agentUpdate = agent_store.Agent{}
	if err := easyjson.Unmarshal(ctx.PostBody(), &agentUpdate); err != nil {
		log.Panic(err.Error())
	}

	// Make sure ID is the same
	agentUpdate.ID = id

	// Update agent
	agent_store.UpdateAgent(agentUpdate)

	ctx.SetStatusCode(http.StatusNoContent)
}

func deleteAgentHandler(ctx *fasthttp.RequestCtx) {
	// Get id from url
	id, err := uuid.Parse(ctx.UserValue("id").(string))
	if err != nil {
		log.Print(err.Error())
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	agent_store.DeleteAgent(id)
}
