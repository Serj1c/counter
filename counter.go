package counter

import (
	"sync"
	"sync/atomic"
)

type Counter interface {
	Get() int32
	Increment(idx int)
}

type MutexCounter struct {
	value int32
	mu    sync.RWMutex
}

func (c *MutexCounter) Get() int32 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func (c *MutexCounter) Increment(_ int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

type AtomicCounter struct {
	value atomic.Int32
	_     [60]byte
}

func (c *AtomicCounter) Get() int32 {
	return c.value.Load()
}

func (c *AtomicCounter) Increment(_ int) {
	_ = c.value.Add(1)
}

type ShardedAtomicCounter struct {
	shards [4]AtomicCounter
}

func (c *ShardedAtomicCounter) Get() int32 {
	var value int32
	for i := 0; i < len(c.shards); i++ {
		value += c.shards[i].Get()
	}
	return value
}

func (c *ShardedAtomicCounter) Increment(idx int) {
	_ = c.shards[idx].value.Add(1)
}
