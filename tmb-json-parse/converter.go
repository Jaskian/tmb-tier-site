package tmbjsonparse

import (
	"strings"

	"github.com/jaskian/tmb-tier-site/shared"
)

const TMB_TIME_FORMAT = "2006-01-02 15:04:05"

type PhaseFile struct {
	Phase int
	File  []byte
}

func ConvertTMBData(jsonData []byte, previousPhaseFiles ...PhaseFile) (shared.TMBData, error) {
	result := shared.TMBData{}

	d, err := unmarshalTMBJson(jsonData)
	if err != nil {
		return result, err
	}

	oldPhaseData := map[int]tmbData{}
	for _, oldPhaseFile := range previousPhaseFiles {
		d, err := unmarshalTMBJson(oldPhaseFile.File)
		if err != nil {
			panic(err)
		}
		oldPhaseData[oldPhaseFile.Phase] = d
	}

	d.mergeTMBData(oldPhaseData)

	return convertToExportData(d)
}

func convertToExportData(d tmbData) (shared.TMBData, error) {
	result := shared.TMBData{}

	for _, c := range d {
		phaseData, err := getPhaseDataFromLoot(c)
		if err != nil {
			return result, err
		}
		outputChar := shared.Character{
			Name:   c.Name,
			Class:  c.Class,
			Spec:   c.Spec,
			Phases: phaseData,
		}

		populateKeyItems(&outputChar)

		result = append(result, outputChar)
	}

	return result, nil
}

type lootKey struct{ phase, slot int }

func getPhaseDataFromLoot(c character) (map[int]shared.PhaseData, error) {
	// storing loot two ways to make calcs easier
	phaseSlotLoot := map[lootKey][]shared.Loot{}
	phaseLoot := phaseSlotLootColl{}

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

	compiledWishlists := compileWishLists(c.Wishlisted)

	result := map[int]shared.PhaseData{}

	for _, phase := range shared.PHASES {
		result[phase] = shared.PhaseData{
			Wishlist: compiledWishlists[phase],
			Slots:    map[int]shared.SlotData{},
		}
		for _, slot := range shared.SLOTS {
			// we need all the phase loot to calculate in-tier, not just items for that slot
			key := lootKey{phase, int(slot)}
			result[key.phase].Slots[key.slot] = calculateInTier(c, key.slot, phase, phaseSlotLoot[key], phaseLoot)
		}
	}

	return result, nil
}

func compileWishLists(loot []loot) map[int]shared.Wishlist {

	result := map[int]shared.Wishlist{}
	lootByPhase := map[int][]shared.WishlistLoot{}

	for _, i := range loot {
		phaseNum := shared.PhaseMappingInstance[i.InstanceID]

		wlLoot := shared.WishlistLoot{
			Loot:     NewLoot(i, 0, 0),
			Received: i.Pivot.ReceivedWLItem == 1,
		}

		lootByPhase[phaseNum] = append(lootByPhase[phaseNum], wlLoot)
	}

	for phase, wlLoots := range lootByPhase {
		received := 0
		for _, i := range wlLoots {
			if i.Received {
				received++
			}
		}

		result[phase] = shared.Wishlist{
			Received:     received,
			Total:        len(wlLoots),
			WishlistLoot: wlLoots,
		}
	}

	return result
}

type phaseSlotLootColl map[int]map[int][]*shared.Loot

func (k phaseSlotLootColl) AddItem(phase int, slot int, loot *shared.Loot) {
	if k[phase] == nil {
		k[phase] = map[int][]*shared.Loot{}
	}
	k[phase][slot] = append(k[phase][slot], loot)
}

func (k phaseSlotLootColl) GetItems(phase int, slot int) []*shared.Loot {
	return k[phase][slot]
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

func populateKeyItems(c *shared.Character) {
	for _, i := range c.Phases[3].Slots[0].Items {
		if i.ItemID == 47242 {
			c.KeyItems.Trophies++
		}
	}
}
