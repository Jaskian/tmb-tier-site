package tmbjsonparse

import (
	"time"

	"github.com/jaskian/tmb-tier-site/shared"
)

var CutOffDate = time.Date(2023, 06, 22, 0, 0, 0, 0, time.Local)

func calculateInTier(c character, slot int, phase int, slotItems []shared.Loot, allItemsByPhase keyedLootCollection) bool {

	inTier := len(slotItems) > 0

	// extra approve logic
	if !inTier {
		allEligibleItems := getEligibleLoot(phase, slotItems, allItemsByPhase)
		if slot == int(shared.TwoHander) {
			// discount on 2h when you've bought a 1h + OH
			has1h := Any(allEligibleItems, isLootForSlotFunc(int(shared.OneHander)))
			hasOH := Any(allEligibleItems, isLootForSlotFunc(int(shared.Offhand)))
			inTier = has1h && hasOH
		} else if slot == int(shared.OneHander) || slot == int(shared.Offhand) {
			// discount on 1h/OH when you've bought a 2h
			inTier = Any(allEligibleItems, isLootForSlotFunc(int(shared.TwoHander)))
		} else if len(allEligibleItems) > 0 {
			inTier = Any(allEligibleItems, isLootForSlotFunc(slot))
		}
	}

	// deny logic
	if inTier {
		if slot == int(shared.TwoHander) {
			// fury warriors have to buy 2x 2h to get in-tier on 2h
			if c.Class == shared.Warrior && c.Spec == "Fury" {
				inTier = len(slotItems) >= 2
			}
		} else if slot == int(shared.Ring) || slot == int(shared.Trinket) {
			// gotta have 2 rings/trinkets to be eligible
			inTier = len(slotItems) > 1
		}
	}

	// by default, we have items for the slot so it in in-tier
	return inTier
}

func getEligibleLoot(phase int, slotItems []shared.Loot, itemsByPhase keyedLootCollection) []*shared.Loot {
	items := []*shared.Loot{}

	for _, i := range slotItems {
		items = append(items, &i)
	}

	// in p3 we allow 252 p2 loot received in p3 to be eligible for in-tier
	if phase == 3 {
		for _, is := range itemsByPhase[2] {
			for _, item := range is {
				if item.Date.After(CutOffDate) && shared.ULDUAR_ITEMLEVELS[item.ItemID] >= 252 {
					items = append(items, item)
				}
			}
		}
	}

	// need all current phase items
	for _, is := range itemsByPhase[phase] {
		items = append(items, is...)
	}

	return items
}

func isLootForSlotFunc(slot int) func(*shared.Loot) bool {
	return func(l *shared.Loot) bool {
		return l.Slot == slot
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
