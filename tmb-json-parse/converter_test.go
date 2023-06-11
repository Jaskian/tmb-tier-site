package tmbjsonparse

import (
	"testing"

	"github.com/jaskian/tmb-tier-site/shared"
)

func TestConvertTMBJson(t *testing.T) {

	t.Run("In P1, in-tier for 1h/2h/OH after buying 2h weapons", func(t *testing.T) {
		input := buildTestDataWithLoot("Druid", "Feral", shared.TwoHander, shared.Naxx)

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
		input := buildTestDataWithLoot("Warrior", "Fury", shared.TwoHander, shared.Ulduar)

		got, err := convertToExportData(input)
		assertNoError(t, err)

		assertNotInTierInPhase(t, 2, shared.TwoHander, got[0])
		assertInTierInPhase(t, 2, shared.OneHander, got[0])
	})

	t.Run("In P2, Warrior in-tier for 2h after 2 bought", func(t *testing.T) {
		input := buildTestDataWithLoot("Warrior", "Fury", shared.TwoHander, shared.Ulduar)

		input[0].ReceivedLoot = append(input[0].ReceivedLoot, loot{
			InventoryType: int(shared.TwoHander),
			InstanceID:    int(shared.Ulduar),
		})

		got, err := convertToExportData(input)
		assertNoError(t, err)
		assertInTierInPhase(t, 2, shared.TwoHander, got[0])
	})
}

func assertInTierInPhase(t *testing.T, phase int, slot shared.Slot, c Character) {
	t.Helper()
	if !c.Phases[phase][int(slot)].InTier {
		t.Errorf("P%d %v should be In-Tier", phase, slot)
	}
}
func assertNotInTierInPhase(t *testing.T, phase int, slot shared.Slot, c Character) {
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

func buildTestDataWithLoot(class, spec string, slot shared.Slot, instance shared.Instance) tmbData {
	return tmbData{character{
		Class: class,
		Spec:  spec,
		ReceivedLoot: []loot{
			{
				InventoryType: int(slot),
				InstanceID:    int(instance),
			},
		},
	},
	}
}
