package eventlistener

// EventListener base class
type EventListener struct {
	listeners map[string][]chan interface{}
	response  chan interface{}
}

// New event creation
func (e *EventListener) New() *EventListener {
	return &EventListener{
		listeners: make(map[string][]chan interface{}),
		response:  make(chan interface{}),
	}
}

// Attach an event listener
func (e *EventListener) Attach(eventid string, ch chan interface{}) {
	if e.listeners == nil {
		e.listeners = make(map[string][]chan interface{})
	}

	e.listeners[eventid] = append(e.listeners[eventid], ch)

}

// Remove an event listener
func (e *EventListener) Remove(eventid string, ch chan interface{}) {
	if _, ok := e.listeners[eventid]; ok {
		for i := range e.listeners[eventid] {
			if e.listeners[eventid][i] == ch {
				e.listeners[eventid] = append(e.listeners[eventid][:i], e.listeners[eventid][i+1:]...)
				break
			}
		}
	}
}

// Emit emits an event on the Dog struct instance
func (e *EventListener) Emit(eventid string, response interface{}) {
	if _, ok := e.listeners[eventid]; ok {
		for _, handler := range e.listeners[eventid] {
			go func(handler chan interface{}) {
				handler <- response
			}(handler)
		}
	}
}
