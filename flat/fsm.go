package flat

import "fmt"

type State string
type Event string

type StateEvent struct {
	State State
	Event Event
}

type Transition struct {
	NextState State
	Rules     []func() bool
}

type FiniteStateMachine struct {
	CurrentState State
	Table        map[StateEvent]Transition
}

func NewFSM(initial State, table map[StateEvent]Transition) *FiniteStateMachine {
	return &FiniteStateMachine{
		CurrentState: initial,
		Table:        table,
	}
}

func (fsm *FiniteStateMachine) Compute(event Event) {
	fmt.Printf("Processing event: %s\n", event)

	key := StateEvent{State: fsm.CurrentState, Event: event}

	transition, exists := fsm.Table[key]
	if !exists {
		fmt.Printf("  -> Ignored: Nothing happens in %s when %s occurs\n", fsm.CurrentState, event)
		return
	}

	canTransition := true
	for _, ruleFn := range transition.Rules {
		if !ruleFn() {
			canTransition = false
			break
		}
	}

	if canTransition {
		fmt.Printf("  -> Success: Transitioning %s to %s\n", fsm.CurrentState, transition.NextState)
		fsm.CurrentState = transition.NextState
	} else {
		fmt.Printf("  -> Blocked: Rule prevented transition\n")
	}
}
