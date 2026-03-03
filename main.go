package main

import (
	"fmt"

	"github.com/rodrigocitadin/fsm/flat"
	"github.com/rodrigocitadin/fsm/graph"
)

type Event string

const (
	EventCoin Event = "Coin"
	EventPush Event = "Push"
)

func main() {
	sequenceOfEvents := []Event{
		EventPush, // Try to push while locked
		EventCoin, // Insert coin
		EventCoin, // Insert another coin (eats it)
		EventPush, // Walk through
		EventPush, // Try to walk through again
	}

	fmt.Println("--- GRAPH VERSION ---")

	graphVersion(sequenceOfEvents)

	fmt.Println("--- END OF GRAPH VERSION ---\n")
	fmt.Println("--- FLAT VERSION ---")

	flatVersion(sequenceOfEvents)

	fmt.Println("--- END OF FLAT VERSION ---")
}

func graphVersion(sequenceOfEvents []Event) {
	locked := &graph.State{Value: "Locked"}
	unlocked := &graph.State{Value: "Unlocked"}

	isValidCoin := func() bool {
		return true
	}

	// If Locked:
	// - Coin unlocks it
	// - Push does nothing (stays locked)
	locked.Arrows = []graph.Arrow{
		{Edge: unlocked, Event: graph.Event(EventCoin), Rules: []func() bool{isValidCoin}},
		{Edge: locked, Event: graph.Event(EventPush)},
	}

	// If Unlocked:
	// - Push locks it (person walks through)
	// - Coin does nothing (already unlocked, eats your coin)
	unlocked.Arrows = []graph.Arrow{
		{Edge: locked, Event: graph.Event(EventPush)},
		{Edge: unlocked, Event: graph.Event(EventCoin)},
	}

	fsm := graph.NewFSM(locked)

	for _, event := range sequenceOfEvents {
		fsm.Compute(graph.Event(event))
	}
}

func flatVersion(sequenceOfEvents []Event) {
	const (
		StateLocked   flat.State = "Locked"
		StateUnlocked flat.State = "Unlocked"
	)

	// The entire logic of the FSM is defined in one clean, readable block
	transitionTable := map[flat.StateEvent]flat.Transition{
		{State: StateLocked, Event: flat.Event(EventCoin)}:   {NextState: StateUnlocked, Rules: []func() bool{isValidCoin}},
		{State: StateLocked, Event: flat.Event(EventPush)}:   {NextState: StateLocked},
		{State: StateUnlocked, Event: flat.Event(EventPush)}: {NextState: StateLocked},
		{State: StateUnlocked, Event: flat.Event(EventCoin)}: {NextState: StateUnlocked},
	}

	fsm := flat.NewFSM(StateLocked, transitionTable)

	for _, event := range sequenceOfEvents {
		fsm.Compute(flat.Event(event))
	}
}

func isValidCoin() bool {
	return true
}
