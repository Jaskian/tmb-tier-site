package tmbjsonparse

import (
	"time"

	"github.com/jaskian/tmb-tier-site/shared"
)

var CutOffDate = time.Date(2023, 06, 22, 0, 0, 0, 0, time.Local)

func calculateInTier(c character, slot int, phase int, slotItems []shared.Loot, allItemsByPhase keyedLootCollection) shared.SlotData {

	slotData := shared.SlotData{
		InTier: len(slotItems) > 0,
		Items:  slotItems,
	}

	// extra approve logic
	if !slotData.InTier {
		allEligibleItems := getEligibleLoot(phase, slotItems, allItemsByPhase)
		if slot == int(shared.TwoHander) {
			// discount on 2h when you've bought a 1h + OH
			oneHand := CollectInto(allEligibleItems, &slotData.Items, isLootForSlotFunc(int(shared.OneHander)))
			offHand := CollectInto(allEligibleItems, &slotData.Items, isLootForSlotFunc(int(shared.Offhand)))
			slotData.InTier = len(oneHand) > 0 && len(offHand) > 0
		} else if slot == int(shared.OneHander) || slot == int(shared.Offhand) {
			// discount on 1h/OH when you've bought a 2h
			found := CollectInto(allEligibleItems, &slotData.Items, isLootForSlotFunc(int(shared.TwoHander)))
			slotData.InTier = len(found) > 0
		} else if len(allEligibleItems) > 0 {
			found := CollectInto(allEligibleItems, &slotData.Items, isLootForSlotFunc(slot))
			slotData.InTier = len(found) > 0
		}
	}

	// deny logic
	if slotData.InTier {
		if slot == int(shared.TwoHander) {
			// fury warriors have to buy 2x 2h to get in-tier on 2h
			if c.Class == shared.Warrior && c.Spec == "Fury" {
				slotData.InTier = len(slotItems) >= 2
			}
		} else if slot == int(shared.Ring) || slot == int(shared.Trinket) {
			// gotta have 2 rings/trinkets to be eligible
			slotData.InTier = len(slotItems) > 1
		}
	}

	// by default, we have items for the slot so it in in-tier
	return slotData
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

func CollectInto[T any](ts []*T, r *[]T, pred func(*T) bool) (result []*T) {
	for _, t := range ts {
		if pred(t) {
			result = append(result, t)
			*r = append(*r, *t)
		}
	}
	return
}

func Any[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if pred(t) {
			return true
		}
	}
	return false
}
