/*--------------------------------------------------------------
 *  Copyright (c) Albert Yakupov. All rights reserved.
 *  Licensed under the Apache 2 License.
 * title: pubsub.go
 * date: 2025-12-28 10:52:35
 *
 * Use Strict!!!
 *-------------------------------------------------------------*/

package pattern

import (
	"sync"
)

type EventMessage struct {
	Topic string
	Value string
}

// Event for Subscribers
type EventChannel struct {
	Pool sync.Map
}

func (ec *EventChannel) SetEvent(msg EventMessage) {
	ec.Pool.Store(msg.Topic, msg.Value)
}

func (ec EventChannel) GetValue(key string) string {
	if t, ok := ec.Pool.Load(key); ok {
		return t.(string)
	} else {
		return ""
	}
}

// pattern Publisher - Subscriber
type Publisher interface {
	Publish()
	// Subscribe return event channel
	Subscribe() *EventChannel
}
