package db

import (
	"github.com/google/uuid"
	"sync"
)

var database Database

type Database struct {
	sync.RWMutex
	internal map[uuid.UUID]Agent
}

func NewDatabase(size int) *Database {
	return &Database{
		internal: make(map[uuid.UUID]Agent, size),
	}
}

func (rm *Database) Load(key uuid.UUID) (value Agent, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *Database) Delete(key uuid.UUID) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *Database) Store(key uuid.UUID, value Agent) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}

func (rm *Database) Values() []Agent {
	rm.RLock()

	agents := make([]Agent, len(rm.internal))

	for _, v := range rm.internal {
		agents = append(agents, v)
	}

	rm.RUnlock()

	return agents
}

func StartDB() {
	database = *NewDatabase(5)
}
