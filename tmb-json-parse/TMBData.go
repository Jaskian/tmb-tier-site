package tmbjsonparse

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

func NewLoot(i loot, phase int, slot int) Loot {
	t := time.Unix(0, 0)

	if i.Pivot.Date != "" {
		t, _ = time.Parse(TMB_TIME_FORMAT, i.Pivot.Date)
	}

	result := Loot{
		ItemID:   i.ItemID,
		ItemName: i.ItemName,
		Phase:    phase,
		Slot:     slot,
		Date:     t,
	}

	return result
}
