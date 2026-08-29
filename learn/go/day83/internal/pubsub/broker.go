package pubsub

import "sync"

// Broker routes published messages to topic subscribers via channels.
type Broker struct {
	mu   sync.RWMutex
	subs map[string][]chan string
}

func NewBroker() *Broker {
	return &Broker{subs: make(map[string][]chan string)}
}

func (b *Broker) Subscribe(topic string) <-chan string {
	ch := make(chan string, 8)
	b.mu.Lock()
	b.subs[topic] = append(b.subs[topic], ch)
	b.mu.Unlock()
	return ch
}

func (b *Broker) Publish(topic, msg string) int {
	b.mu.RLock()
	targets := append([]chan string(nil), b.subs[topic]...)
	b.mu.RUnlock()
	delivered := 0
	for _, ch := range targets {
		select {
		case ch <- msg:
			delivered++
		default:
		}
	}
	return delivered
}

func (b *Broker) Close(topic string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, ch := range b.subs[topic] {
		close(ch)
	}
	delete(b.subs, topic)
}
