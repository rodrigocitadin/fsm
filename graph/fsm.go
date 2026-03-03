package graph

import "fmt"

type Event string

type State struct {
	Value  string
	Arrows []Arrow
}

type Arrow struct {
	Edge  *State
	Event Event
	Rules []func() bool
}

type FiniteStateMachine struct {
	CurrentState *State
}

func NewFSM(initialState *State) *FiniteStateMachine {
	return &FiniteStateMachine{
		CurrentState: initialState,
	}
}

func (fsm *FiniteStateMachine) Compute(events []Event) {
	for _, event := range events {
		fmt.Printf("Processing event: %s\n", event)

		transitioned := false

		for _, arrow := range fsm.CurrentState.Arrows {
			if event == arrow.Event {
				canTransition := true
				for _, ruleFn := range arrow.Rules {
					if !ruleFn() {
						canTransition = false
						break
					}
				}

				if canTransition {
					fmt.Printf("  -> Success: Transitioning %s to %s\n", fsm.CurrentState.Value, arrow.Edge.Value)
					fsm.CurrentState = arrow.Edge
					transitioned = true
					break
				} else {
					fmt.Printf("  -> Blocked: Rule prevented transition\n")
				}
			}
		}

		if !transitioned {
			fmt.Printf("  -> Ignored: Nothing happens in %s when %s occurs\n", fsm.CurrentState.Value, event)
		}
	}
}
