package db

import (
	"errors"
	"github.com/google/uuid"
)

// Agent defines defines and an agent that's part of the simulation
type Agent struct {
	ID      uuid.UUID `json:"id"`
	X, Y, Z float64
}

// NewAgent creates a new agent
func NewAgent(a Agent) (Agent, error) {
	id, _ := uuid.NewRandom()
	a.ID = id

	database.Store(id, a)

	return a, nil
}

// GetAgent return the agent with the specified UUID
func GetAgent(id uuid.UUID) (Agent, error) {
	agent, ok := database.Load(id)
	if !ok {
		return Agent{}, errors.New("Agent not present")
	}
	return agent, nil
}

// UpdateAgent updates the state of an agent
func UpdateAgent(agent Agent) error {
	database.Store(agent.ID, agent)

	return nil
}

// DeleteAgent removes agent from db
func DeleteAgent(id uuid.UUID) error {
	database.Delete(id)
	return nil
}

// GetAllAgents returns every agent registered to the controller
func GetAllAgents() ([]Agent, error) {

	return database.Values(), nil
}
