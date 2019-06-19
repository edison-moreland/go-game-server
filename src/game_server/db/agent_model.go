package db

import (
	"fmt"
	"github.com/google/uuid"
)

// Agent defines defines and an agent that's part of the simulation
type Agent struct {
	ID      uuid.UUID `gorm:"unique;not null" json:"id"`
	X, Y, Z float64
}

// NewAgent creates a new agent
func NewAgent(a Agent) (Agent, error) {
	dbConn, err := DBConnection()
	if err != nil {
		return Agent{}, err
	}

	id, _ := uuid.NewRandom()
	a.ID = id

	dbConn.Create(&a)

	return a, nil
}

// GetAgent return the agent with the specified UUID
func GetAgent(uuid uuid.UUID) (Agent, error) {
	dbConn, err := DBConnection()
	if err != nil {
		return Agent{}, err
	}

	agent := Agent{}
	dbConn.Where(&Agent{ID: uuid}).First(&agent)

	if agent.ID != uuid {
		return agent, fmt.Errorf("Couldn't find agent with uuid '%v'", uuid)
	}

	return agent, nil
}

// UpdateAgent updates the state of an agent
func UpdateAgent(agent Agent) error {
	dbConn, err := DBConnection()
	if err != nil {
		return err
	}

	dbConn.Save(&agent)

	return nil
}

// DeleteAgent removes agent from db
func DeleteAgent(id uuid.UUID) error {
	dbConn, err := DBConnection()
	if err != nil {
		return err
	}

	agent, err := GetAgent(id)
	if err != nil {
		return err
	}

	dbConn.Delete(&agent)

	return nil
}

// GetAllAgents returns every agent registered to the controller
func GetAllAgents() ([]Agent, error) {
	dbConn, err := DBConnection()
	if err != nil {
		return nil, err
	}

	agents := []Agent{}
	dbConn.Find(&agents)
	return agents, nil
}
