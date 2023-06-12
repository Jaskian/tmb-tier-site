package tmbjsonparse

import (
	"testing"

	"github.com/jaskian/tmb-tier-site/shared"
)

func TestConvertTMBJson(t *testing.T) {

	t.Run("In P1, in-tier for 1h/2h/OH after buying 2h weapons", func(t *testing.T) {
		input := buildTestDataWithLoot("Druid", "Feral", shared.TwoHander, shared.Naxx, 0)

		got, err := convertToExportData(input)
		assertNoError(t, err)

		assertInTierInPhase(t, 1, shared.TwoHander, got[0])
		assertInTierInPhase(t, 1, shared.OneHander, got[0])
		assertInTierInPhase(t, 1, shared.Offhand, got[0])
		// spot checks
		assertNotInTierInPhase(t, 2, shared.TwoHander, got[0])
		assertNotInTierInPhase(t, 1, shared.Neck, got[0])
	})

	t.Run("In P2, Warrior no in-tier for 2h with only 1 bought", func(t *testing.T) {
		input := buildTestDataWithLoot("Warrior", "Fury", shared.TwoHander, shared.Ulduar, 0)

		got, err := convertToExportData(input)
		assertNoError(t, err)

		assertNotInTierInPhase(t, 2, shared.TwoHander, got[0])
		assertInTierInPhase(t, 2, shared.OneHander, got[0])
	})

	t.Run("In P2, Warrior in-tier for 2h after 2 bought", func(t *testing.T) {
		input := buildTestDataWithLoot("Warrior", "Fury", shared.TwoHander, shared.Ulduar, 0)

		input[0].ReceivedLoot = append(input[0].ReceivedLoot, loot{
			InventoryType: int(shared.TwoHander),
			InstanceID:    int(shared.Ulduar),
		})

		got, err := convertToExportData(input)
		assertNoError(t, err)
		assertInTierInPhase(t, 2, shared.TwoHander, got[0])
	})

	t.Run("In P2, no in-tier for trinket when 1 bought", func(t *testing.T) {
		input := buildTestDataWithLoot("Warrior", "Fury", shared.Trinket, shared.Ulduar, 0)

		got, err := convertToExportData(input)
		assertNoError(t, err)
		assertNotInTierInPhase(t, 2, shared.Ring, got[0])
	})

	t.Run("In P2, in-tier for rings when 2 bought", func(t *testing.T) {
		input := buildTestDataWithLoot("Warrior", "Fury", shared.Ring, shared.Ulduar, 0)

		input[0].ReceivedLoot = append(input[0].ReceivedLoot, loot{
			InventoryType: int(shared.Ring),
			InstanceID:    int(shared.Ulduar),
		})

		got, err := convertToExportData(input)
		assertNoError(t, err)
		assertInTierInPhase(t, 2, shared.Ring, got[0])
	})

	t.Run("Offspec items excluded", func(t *testing.T) {
		input := buildTestDataWithLoot("Warrior", "Fury", shared.Belt, shared.Ulduar, 1)

		offspecTexts := []string{"~Off-Spec~", "~Banking~", "~Free~"}
		for _, t := range offspecTexts {
			input[0].ReceivedLoot = append(input[0].ReceivedLoot, loot{
				InventoryType: int(shared.Belt),
				InstanceID:    int(shared.Ulduar),
				Pivot:         pivot{OfficerNote: t},
			})
		}

		got, err := convertToExportData(input)
		assertNoError(t, err)
		assertNotInTierInPhase(t, 2, shared.Belt, got[0])

		count := len(got[0].Phases[2][int(shared.Belt)].Items)

		if count > 0 {
			t.Errorf("Expected 0 items, got %d", count)
		}
	})
}

func assertInTierInPhase(t *testing.T, phase int, slot shared.Slot, c shared.Character) {
	t.Helper()
	if !c.Phases[phase][int(slot)].InTier {
		t.Errorf("P%d %v should be In-Tier", phase, slot)
	}
}
func assertNotInTierInPhase(t *testing.T, phase int, slot shared.Slot, c shared.Character) {
	t.Helper()
	if c.Phases[phase][int(slot)].InTier {
		t.Errorf("P%d %v should not be In-Tier", phase, slot)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func buildTestDataWithLoot(class, spec string, slot shared.Slot, instance shared.Instance, offspec int) tmbData {

	return tmbData{character{
		Class: class,
		Spec:  spec,
		ReceivedLoot: []loot{
			{
				InventoryType: int(slot),
				InstanceID:    int(instance),
				Pivot:         pivot{Offspec: offspec},
			},
		},
	},
	}
}
