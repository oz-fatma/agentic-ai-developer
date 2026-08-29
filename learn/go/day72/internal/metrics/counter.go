package metrics

import (
	"fmt"
	"sync/atomic"
)

type Counter struct {
	name  string
	value atomic.Int64
}

func NewCounter(name string) *Counter {
	return &Counter{name: name}
}

func (c *Counter) Inc() {
	c.value.Add(1)
}

func (c *Counter) Add(n int64) {
	c.value.Add(n)
}

func (c *Counter) Value() int64 {
	return c.value.Load()
}

func (c *Counter) String() string {
	return fmt.Sprintf("%s=%d", c.name, c.Value())
}
