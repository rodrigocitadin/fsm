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

const (
	EventCoin Event = "Coin"
	EventPush Event = "Push"
)

func main() {
	locked := &State{Name: "Locked"}
	unlocked := &State{Name: "Unlocked"}

	isValidCoin := func() bool {
		return true
	}

	// If Locked:
	// - Coin unlocks it
	// - Push does nothing (stays locked)
	locked.Arrows = []Arrow{
		{Edge: unlocked, Event: EventCoin, Rules: []func() bool{isValidCoin}},
		{Edge: locked, Event: EventPush},
	}

	// If Unlocked:
	// - Push locks it (person walks through)
	// - Coin does nothing (already unlocked, eats your coin)
	unlocked.Arrows = []Arrow{
		{Edge: locked, Event: EventPush},
		{Edge: unlocked, Event: EventCoin},
	}

	fsm := NewFiniteStateMachine(locked)

	fmt.Printf("Initial State: %s\n\n", fsm.CurrentState.Name)

	sequenceOfEvents := []Event{
		EventPush, // Try to push while locked
		EventCoin, // Insert coin
		EventCoin, // Insert another coin (eats it)
		EventPush, // Walk through
		EventPush, // Try to walk through again
	}

	fsm.Compute(sequenceOfEvents)

	fmt.Printf("Final State: %s\n", fsm.CurrentState.Name)
}
