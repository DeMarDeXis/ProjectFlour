package events

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type EventHandler func(event Event)

type EventBus struct {
	subscribers map[string]map[string]EventHandler
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string]map[string]EventHandler),
	}
}

// Subscribe add user to event's topic and return ID
func (eb *EventBus) Subscribe(topic string, handler EventHandler) string {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if _, exists := eb.subscribers[topic]; !exists {
		eb.subscribers[topic] = make(map[string]EventHandler)
	}

	//id := generateUniqueID()
	id := "sub_" + fmt.Sprintf("%d", time.Now().UnixNano())
	//id := "sub_" + uuid.New().String() // TODO: It's maybe better to use uuid
	eb.subscribers[topic][id] = handler
	return id
}

// Unsubscribe using ID
func (eb *EventBus) Unsubscribe(topic string, id string) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if subs, exists := eb.subscribers[topic]; exists {
		delete(subs, id)
		if len(subs) == 0 {
			delete(eb.subscribers, topic)
		}
	}
}

// Publish send data to all subscribers of the specified topic
func (eb *EventBus) Publish(topic string, event Event) error {
	const op = "internal.events.eventBus.Publish"

	if err := event.Validate(); err != nil {
		return fmt.Errorf("%s.Validate: %w", op, err)
	}

	eb.mu.RLock()
	defer eb.mu.RUnlock()

	if handlers, exists := eb.subscribers[topic]; exists {
		for _, handler := range handlers {
			go func(h EventHandler) {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("Recovered from panic: %v, in %s", r, op)
					}
				}()
				h(event)
			}(handler)
		}
	}

	return nil
}

// generateUniqueID subfunc for generate unique ID
func generateUniqueID() string {
	return "sub_" + fmt.Sprintf("%d", time.Now().UnixNano())
}

// TODO: future
//func (eb *EventBus) SubscribeWithContext(ctx context.Context, topic string, handler func(interface{})) string {
//	id := eb.Subscribe(topic, handler)
//
//	go func() {
//		<-ctx.Done()
//		eb.Unsubscribe(topic, id)
//	}()
//
//	return id
//}
