package tmbjsonparse

import (
	"strings"

	"github.com/jaskian/tmb-tier-site/shared"
)

const TMB_TIME_FORMAT = "2006-01-02 15:04:05"

func ConvertTMBData(jsonData []byte) (shared.TMBData, error) {
	result := shared.TMBData{}

	d, err := unmarshalTMBJson(jsonData)
	if err != nil {
		return result, err
	}

	return convertToExportData(d)
}

func convertToExportData(d tmbData) (shared.TMBData, error) {
	result := shared.TMBData{}

	for _, c := range d {
		phaseData, err := getPhaseDataFromLoot(c)
		if err != nil {
			return result, err
		}
		wlItems := getWishlistDataFromCharacter(c)

		outputChar := shared.Character{
			Name:     c.Name,
			Class:    c.Class,
			Spec:     c.Spec,
			Phases:   phaseData,
			Wishlist: wlItems,
		}
		result = append(result, outputChar)
	}

	return result, nil
}

func getWishlistDataFromCharacter(c character) shared.Wishlist {

	loot := []shared.WishlistLoot{}
	received := 0

	for _, i := range c.Wishlisted {
		li := shared.WishlistLoot{
			Loot:     NewLoot(i, 0, 0),
			Received: i.Pivot.ReceivedWLItem == 1,
		}
		if li.Received {
			received++
		}
		loot = append(loot, li)
	}

	wl := shared.Wishlist{
		Total:        len(loot),
		Received:     received,
		WishlistLoot: loot,
	}

	return wl
}

type lootKey struct{ phase, slot int }

func getPhaseDataFromLoot(c character) (map[int]shared.PhaseData, error) {
	// storing loot two ways to make calcs easier
	phaseSlotLoot := map[lootKey][]shared.Loot{}
	phaseLoot := keyedLoot{}

	for _, i := range c.ReceivedLoot {

		if i.ExcludedFromResults() {
			continue
		}

		phaseNum := shared.PhaseMappingInstance[i.InstanceID]

		slotNum := i.GetSlot()

		loot := NewLoot(i, phaseNum, slotNum)
		key := lootKey{phaseNum, slotNum}
		phaseSlotLoot[key] = append(phaseSlotLoot[key], loot)
		phaseLoot.AddItem(key.phase, key.slot, &loot)
	}

	result := newPhasesData()
	for _, phase := range shared.PHASES {
		for _, slot := range shared.SLOTS {
			// we need all the phase loot to calculate in-tier, not just items for that slot
			slotNum := int(slot)
			key := lootKey{phase, slotNum}
			inTier := calculateInTier(c, slotNum, phase, phaseSlotLoot[key], phaseLoot)
			result[key.phase][key.slot] = shared.SlotData{InTier: inTier, Items: phaseSlotLoot[key]}
		}
	}

	return result, nil
}

type keyedLoot map[int]map[int][]*shared.Loot

func (k keyedLoot) AddItem(phase int, slot int, loot *shared.Loot) {
	if k[phase] == nil {
		k[phase] = map[int][]*shared.Loot{}
	}
	k[phase][slot] = append(k[phase][slot], loot)
}

func (k keyedLoot) GetItems(phase int, slot int) []*shared.Loot {
	return k[phase][slot]
}

func newPhasesData() map[int]shared.PhaseData {
	phasesData := map[int]shared.PhaseData{}

	for _, p := range shared.PHASES {
		phasesData[p] = shared.PhaseData{}
	}

	return phasesData
}

func (l loot) ExcludedFromResults() bool {
	if l.Pivot.Offspec == 1 ||
		strings.Contains(l.Pivot.OfficerNote, "Off-Spec") ||
		strings.Contains(l.Pivot.OfficerNote, "Banking") ||
		strings.Contains(l.Pivot.OfficerNote, "Free") {
		return true
	}
	return false
}

func (l loot) GetSlot() int {
	slotNum := l.InventoryType

	if slotNum == 0 {
		firstWord := strings.Split(l.ItemName, " ")[0]
		if slot, ok := shared.TokenMapping[firstWord]; ok {
			slotNum = int(slot)
		}
	} else if slot, ok := shared.InventoryTypeMappings[slotNum]; ok {
		slotNum = int(slot)
	}

	return slotNum
}
