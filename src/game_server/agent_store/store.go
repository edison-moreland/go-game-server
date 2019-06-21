package agent_store

import (
	"github.com/google/uuid"
	"sync"
)

type agentStore struct {
	sync.RWMutex
	internal map[uuid.UUID]Agent
}

func newAgentStore(size int) *agentStore {
	return &agentStore{
		internal: make(map[uuid.UUID]Agent, size),
	}
}

func (rm *agentStore) Load(key uuid.UUID) (value Agent, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *agentStore) Delete(key uuid.UUID) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *agentStore) Store(key uuid.UUID, value Agent) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}

func (rm *agentStore) Values() []Agent {
	rm.RLock()

	agents := make([]Agent, len(rm.internal))

	for _, v := range rm.internal {
		agents = append(agents, v)
	}

	rm.RUnlock()

	return agents
}
