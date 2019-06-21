package agent_store

import (
	"errors"
	"github.com/google/uuid"
)

var agents agentStore

func init() {
	agents = *newAgentStore(5)
}

// Agent defines defines and an agent that's part of the simulation
//easyjson:json
type Agent struct {
	ID      uuid.UUID `json:"id"`
	X, Y, Z float64
}

// AgentList is used instead of a raw []Agent for easier unmarshal
//easyjson:json
type AgentList struct {
	Agents []Agent `json:"agents"`
}

// NewAgent creates a new agent
func NewAgent(a Agent) (Agent, error) {
	id, _ := uuid.NewRandom()
	a.ID = id

	agents.Store(id, a)

	return a, nil
}

// GetAgent return the agent with the specified UUID
func GetAgent(id uuid.UUID) (Agent, error) {
	agent, ok := agents.Load(id)
	if !ok {
		return Agent{}, errors.New("Agent not present")
	}
	return agent, nil
}

// UpdateAgent updates the state of an agent
func UpdateAgent(agent Agent) error {
	agents.Store(agent.ID, agent)
	return nil
}

// DeleteAgent removes agent from db
func DeleteAgent(id uuid.UUID) error {
	agents.Delete(id)
	return nil
}

// GetAllAgents returns every agent registered to the controller
func GetAllAgents() (AgentList, error) {
	return AgentList{agents.Values()}, nil
}
