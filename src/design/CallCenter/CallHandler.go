package CallCenter

import (
	"fmt"
)

type Rank int

const (
	RESPONDER Rank = 1 + iota
	MANAGER
	DIRECTOR
	NUM_RESPONDENTS = 10
	NUM_MANAGERS = 4
	NUM_DIRECTORS = 2
	LEVELS = 3
)

type CallHandler struct {
	employeeLevels [][]Employee
	callQueues [][]Call
}

func newHandler() *CallHandler {
	return &CallHandler{
	}
}

func (ch *CallHandler) getHandlerForCall(call Call){
	for e := range ch.employeeLevels[0]{
		if 
	}
}

// from caller to employee
func (ch *CallHandler) dispatchCall(caller Caller){
	call := NewCall(caller)
	ch.dispatchCall(call)
}

func (ch *CallHandler) dispatchCall(call Call){
	emp := get
}