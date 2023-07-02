package tmbjsonparse

import (
	"embed"
	"testing"

	"github.com/jaskian/tmb-tier-site/shared"
)

func TestConvertTMBJson(t *testing.T) {

	t.Run("In P1, in-tier for 1h/2h/OH after buying 2h weapons", func(t *testing.T) {
		input := buildTestDataWithLoot("Druid", "Feral", shared.TwoHander, shared.Naxx, 0, "")

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
		input := buildTestDataWithLoot("Warrior", "Fury", shared.TwoHander, shared.Ulduar, 0, "")

		got, err := convertToExportData(input)
		assertNoError(t, err)

		assertNotInTierInPhase(t, 2, shared.TwoHander, got[0])
		assertInTierInPhase(t, 2, shared.OneHander, got[0])
	})

	t.Run("In P2, Warrior in-tier for 2h after 2 bought", func(t *testing.T) {
		input := buildTestDataWithLoot("Warrior", "Fury", shared.TwoHander, shared.Ulduar, 0, "")

		input[0].ReceivedLoot = append(input[0].ReceivedLoot, loot{
			InventoryType: int(shared.TwoHander),
			InstanceID:    int(shared.Ulduar),
		})

		got, err := convertToExportData(input)
		assertNoError(t, err)
		assertInTierInPhase(t, 2, shared.TwoHander, got[0])
	})

	t.Run("In P2, no in-tier for trinket when 1 bought", func(t *testing.T) {
		input := buildTestDataWithLoot("Warrior", "Fury", shared.Trinket, shared.Ulduar, 0, "")

		got, err := convertToExportData(input)
		assertNoError(t, err)
		assertNotInTierInPhase(t, 2, shared.Ring, got[0])
	})

	t.Run("In P2, in-tier for rings when 2 bought", func(t *testing.T) {
		input := buildTestDataWithLoot("Warrior", "Fury", shared.Ring, shared.Ulduar, 0, "")

		input[0].ReceivedLoot = append(input[0].ReceivedLoot, loot{
			InventoryType: int(shared.Ring),
			InstanceID:    int(shared.Ulduar),
		})

		got, err := convertToExportData(input)
		assertNoError(t, err)
		assertInTierInPhase(t, 2, shared.Ring, got[0])
	})

	t.Run("In P3, P2 252 items contribute to in-tier when bought in P3", func(t *testing.T) {

		p2Date := "2023-05-24 15:04:05"
		p3Date := "2023-06-24 15:04:05"

		input := tmbData{character{
			Class: "Warrior",
			Spec:  "Fury",
			ReceivedLoot: []loot{
				{
					ItemID:        45132, // 252
					InventoryType: int(shared.Belt),
					InstanceID:    int(shared.Ulduar),
					Pivot:         pivot{Date: p2Date},
				},
				{
					ItemID:        45132, // 252
					InventoryType: int(shared.Wrist),
					InstanceID:    int(shared.Ulduar),
					Pivot:         pivot{Date: p3Date},
				},
			},
		},
		}

		got, err := convertToExportData(input)
		assertNoError(t, err)

		// in p2 but not p3 for old loot
		assertInTierInPhase(t, 2, shared.Belt, got[0])
		assertNotInTierInPhase(t, 3, shared.Belt, got[0])
		// in both for new loot
		assertInTierInPhase(t, 2, shared.Wrist, got[0])
		assertInTierInPhase(t, 3, shared.Wrist, got[0])

		p3WristItemCount := len(got[0].Phases[3].Slots[int(shared.Wrist)].Items)
		if p3WristItemCount != 1 {
			t.Errorf("Expected 1, got %d", p3WristItemCount)
		}
	})

	t.Run("In P3, P2 252 items bought in-tier in Ulduar do not count towards P3 in-tier", func(t *testing.T) {

		p3Date := "2023-06-24 15:04:05"

		input := tmbData{character{
			Class: "Warrior",
			Spec:  "Fury",
			ReceivedLoot: []loot{
				{
					ItemID:        45132, // 252
					InventoryType: int(shared.Belt),
					InstanceID:    int(shared.Ulduar),
					Pivot:         pivot{Date: p3Date, OfficerNote: "~In-tier Upgrade~"},
				},
				{
					ItemID:        45132, // 252
					InventoryType: int(shared.Belt),
					InstanceID:    int(shared.Ulduar),
					Pivot:         pivot{Date: p3Date, OfficerNote: "~In tier Upgrade~"},
				},
			},
		},
		}

		got, err := convertToExportData(input)
		assertNoError(t, err)
		assertNotInTierInPhase(t, 3, shared.Belt, got[0])
	})

	t.Run("In P3, TOTC items affect P3 WL", func(t *testing.T) {
		input := buildTestDataWithLoot("Warrior", "Fury", shared.Belt, shared.Totc25, 0, "")
		got, err := convertToExportData(input)
		assertNoError(t, err)
		assertInTierInPhase(t, 3, shared.Belt, got[0])
	})

	t.Run("Trophies are counted for TOTC25", func(t *testing.T) {
		input := tmbData{character{
			ReceivedLoot: []loot{
				{
					ItemID:     47242,
					InstanceID: int(shared.Totc10),
					Pivot:      pivot{Offspec: 0},
				},
			},
		},
		}

		export, err := convertToExportData(input)
		assertNoError(t, err)

		want := 1
		got := export[0].KeyItems.Trophies
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}

	})

	t.Run("Offspec items excluded", func(t *testing.T) {
		input := buildTestDataWithLoot("Warrior", "Fury", shared.Belt, shared.Ulduar, 1, "")

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

		count := len(got[0].Phases[2].Slots[int(shared.Belt)].Items)

		if count > 0 {
			t.Errorf("Expected 0 items, got %d", count)
		}
	})

	t.Run("Wishlist items are sorted into the right phase", func(t *testing.T) {
		data := tmbData{character{
			Wishlisted: []loot{
				loot{InstanceID: 20},
				loot{InstanceID: 28},
				loot{InstanceID: 32},
			},
		},
		}

		got, err := convertToExportData(data)
		assertNoError(t, err)
		for _, phase := range got[0].Phases {
			len := len(phase.Wishlist.WishlistLoot)
			if len != 1 {
				t.Errorf("Expected 1 item, got %d", len)
			}
		}
	})
}

func assertInTierInPhase(t *testing.T, phase int, slot shared.Slot, c shared.Character) {
	t.Helper()
	if !c.Phases[phase].Slots[int(slot)].InTier {
		t.Errorf("P%d %v should be In-Tier", phase, slot)
	}
}
func assertNotInTierInPhase(t *testing.T, phase int, slot shared.Slot, c shared.Character) {
	t.Helper()
	if c.Phases[phase].Slots[int(slot)].InTier {
		t.Errorf("P%d %v should not be In-Tier", phase, slot)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

var (
	//go:embed data/character-json.json
	characterJson []byte
	embedKeeper   embed.FS
)

func BenchmarkConvertTMBData(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ConvertTMBData(characterJson)
	}
}

func buildTestDataWithLoot(class, spec string, slot shared.Slot, instance shared.Instance, offspec int, time string) tmbData {

	return tmbData{character{
		Class: class,
		Spec:  spec,
		ReceivedLoot: []loot{
			{
				InventoryType: int(slot),
				InstanceID:    int(instance),
				Pivot:         pivot{Offspec: offspec, Date: time},
			},
		},
	},
	}
}
