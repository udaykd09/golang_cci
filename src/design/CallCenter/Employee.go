package CallCenter

type Rank int

const (
	RESPONDER Rank = 1 + iota
	MANAGER
	DIRECTOR
)

type Employee struct {
	currentCall Call
	rank Rank
}

func NewEmployee(ch CallHandler) Employee {
}

func (e *Employee) receiveCall(call Call){
	call.reply("Hello")
}

func (e *Employee) callCompleted(call Call){
	call.reply("Thank you!")
	call.disconnect()
}

func (e *Employee) escalate(call Call){
	call.reply("escalating to next level")
	call.increamentRank()
	e.assignNewCall()
}

func (e *Employee) isFree(){
	if e.currentCall != nil {
		return false
	} else {
		return true
	}
}

type Respondent struct {
	Employee
}

type Manager struct {
	Employee
}

type Director struct {
	Employee
}