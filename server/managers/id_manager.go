package managers

import "sync/atomic"

type IdManager struct {
	lastId uint32
}

func NewIdManager() *IdManager {
	return &IdManager{
		lastId: 0,
	}
}

func (i *IdManager) GenerateClientId() uint32 {
	newClientId := atomic.AddUint32(&i.lastId, 1)
	i.lastId = newClientId
	return newClientId
}
