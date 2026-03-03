package main

import "fmt"

type Event string

type State struct {
	Name   string
	Value  any
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

func NewFiniteStateMachine(initialState *State) *FiniteStateMachine {
	return &FiniteStateMachine{
		CurrentState: initialState,
	}
}

func (fsm *FiniteStateMachine) Compute(events []Event) {
	for _, event := range events {
		fmt.Printf("Processing event: %s\n", event)

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
					fmt.Printf("  -> Transitioning from %s to %s\n", fsm.CurrentState.Name, arrow.Edge.Name)
					fsm.CurrentState = arrow.Edge
					break
				} else {
					fmt.Printf("  -> Transition blocked by rule\n")
				}
			}
		}
	}
}
