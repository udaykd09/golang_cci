package CallCenter

import (
	"fmt"
)

type Caller struct {
	name string
}

type Call struct {
	rank Rank
	caller Caller
	handler Employee
	chat []string
}

func NewCall(c Caller) *Call {
	c := call{
		rank: RESPONDER,
		caller: c
	}
	return &c
}

func (c *Call) setHandler(emp Employee) {
	c.handler = emp
}

func (c *Call) printChat() {
	fmt.Println(c.chat)
}

func (c *Call) increamentRank() {
	if c.rank < 2 {
		c.rank += 1
	}
}

func (c *Call) reply(msg string) {
	chat = append(c.chat, msg)
}

func (c *Call) disconnect() {
	fmt.Println("Ending Call...")
	// Free the Employee handling the call
	fmt.Println(c.chat)
}
