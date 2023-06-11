package tmbjsonparse

import (
	"github.com/jaskian/tmb-tier-site/shared"
)

const TMB_TIME_FORMAT = "2006-01-02 15:04:05"

func ConvertTMBData(jsonData []byte) (TMBData, error) {
	result := TMBData{}

	d, err := unmarshalTMBJson(jsonData)
	if err != nil {
		return result, err
	}

	return convertToExportData(d)
}

func convertToExportData(d tmbData) (TMBData, error) {
	result := TMBData{}

	for _, c := range d {
		phaseData, err := getPhaseDataFromLoot(c)
		if err != nil {
			return result, err
		}
		outputChar := Character{
			Name:   c.Name,
			Class:  c.Class,
			Spec:   c.Spec,
			Phases: phaseData,
		}
		result = append(result, outputChar)
	}

	return result, nil
}

type lootKey struct{ phase, slot int }

func getPhaseDataFromLoot(c character) (map[int]PhaseData, error) {
	result := newPhasesData()

	// storing loot two ways to make calcs easier
	phaseSlotLoot := map[lootKey][]Loot{}
	phaseLoot := map[int][]*Loot{}

	for _, i := range c.ReceivedLoot {

		phaseNum := shared.PhaseMappingInstance[i.InstanceID]

		slotNum := i.InventoryType
		if slot, ok := shared.InventoryTypeMappings[slotNum]; ok {
			slotNum = int(slot)
		}

		loot, err := NewLoot(i, phaseNum, slotNum)
		if err != nil {
			return result, err
		}

		key := lootKey{phaseNum, slotNum}
		phaseSlotLoot[key] = append(phaseSlotLoot[key], loot)
		phaseLoot[key.phase] = append(phaseLoot[key.phase], &loot)
	}

	for _, phase := range shared.PHASES {
		for _, slot := range shared.SLOTS {
			// we need all the phase loot to calculate in-tier, not just items for that slot
			slotNum := int(slot)
			key := lootKey{phase, slotNum}
			inTier := calculateInTier(c, slotNum, phaseSlotLoot[key], phaseLoot[phase])
			result[key.phase][key.slot] = SlotData{InTier: inTier, Items: phaseSlotLoot[key]}
		}
	}

	return result, nil
}

func calculateInTier(c character, slot int, slotItems []Loot, allPhaseItems []*Loot) bool {

	inTier := len(slotItems) > 0

	// extra approve logic
	if !inTier {
		if slot == int(shared.TwoHander) {
			// discount on 2h when you've bought a 1h + OH
			has1h := Any(allPhaseItems, isLootForSlotFunc(shared.OneHander))
			hasOH := Any(allPhaseItems, isLootForSlotFunc(shared.Offhand))
			inTier = has1h && hasOH
		} else if slot == int(shared.OneHander) || slot == int(shared.Offhand) {
			// discount on 1h/OH when you've bought a 2h
			inTier = Any(allPhaseItems, isLootForSlotFunc(shared.TwoHander))
		}
	}

	// deny logic
	if inTier {
		if slot == int(shared.TwoHander) {
			// fury warriors have to buy 2x 2h to get in-tier on 2h
			if c.Class == shared.Warrior && c.Spec == "Fury" {
				inTier = len(slotItems) >= 2
			}
		}
	}

	// by default, we have items for the slot so it in in-tier
	return inTier
}

func newPhasesData() map[int]PhaseData {

	phasesData := map[int]PhaseData{}

	for _, p := range shared.PHASES {
		// pd := PhaseData{}
		// for _, s := range shared.SLOTS {
		// 	pd[int(s)] = SlotData{Items: []Loot{}}
		// }
		phasesData[p] = PhaseData{}
	}

	return phasesData
}

func isLootForSlotFunc(slot shared.Slot) func(*Loot) bool {
	return func(l *Loot) bool {
		return l.Slot == int(slot)
	}
}

func Any[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if pred(t) {
			return true
		}
	}
	return false
}
