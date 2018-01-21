package CallCenter

const (
	NUM_RESPONDENTS = 10
	NUM_MANAGERS = 4
	NUM_DIRECTORS = 2
	LEVELS = 3
)

type CallHandler struct {
	employeeLevels [][]Employee
	callQueues [][]Call
}

func NewHandler() *CallHandler {
	return &CallHandler{
	}
}

func addEmployees(rank Rank, e []Employee){
	
}

func (ch *CallHandler) getHandlerForCall(call Call) *Employee{
	for _, e := range ch.employeeLevels[0]{
		if e.isFree() {
			return &e
		} else {
			return nil
		}
	}
}

// Create call obj
func (ch *CallHandler) dispatchCall(caller Caller){
	call := NewCall(caller)
	ch.dispatchCall(call)
}

// find free employee and assign call  
func (ch *CallHandler) dispatchCall(call Call){
	emp := ch.getHandlerForCall(call)
	if emp != nil {
		emp.receiveCall(call)
		call.setHandler(emp)
	} else {
		// Append to call
		call.reply("Please wait in Queue")
		rank := call.getRank()
		ch.callQueues[rank] = append(ch.callQueues[rank], call)
	}
}

func (ch *CallHandler) assignCallFromQueue(e Employee) bool {
		// Assign call to the employee from queue 
		// and remove the call from the queue 
}