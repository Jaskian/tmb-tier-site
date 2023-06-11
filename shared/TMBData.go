package shared

import "time"

type TMBData []Character

type Character struct {
	Name   string
	Class  string
	Spec   string
	Phases map[int]PhaseData
}

type PhaseData map[int]SlotData

type SlotData struct {
	InTier bool
	Items  []Loot
}

type Loot struct {
	ItemID   int
	ItemName string
	Phase    int
	Slot     int
	Date     time.Time
}
