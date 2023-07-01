package webbuilder

import (
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/jaskian/tmb-tier-site/shared"
)

func TestBuildWebsite(t *testing.T) {

	phaseData := createPhaseData()
	insertSlotData(&phaseData, true, shared.Loot{
		ItemName: "Thunderfury",
		Slot:     int(shared.TwoHander),
		Phase:    2,
	})

	data := shared.TMBData{
		shared.Character{
			Name:   "Jaskia",
			Class:  "Hunter",
			Spec:   "Survival",
			Phases: phaseData,
		},
		shared.Character{
			Name:   "Youngart",
			Class:  "Hunter",
			Spec:   "Survival",
			Phases: phaseData,
		},
	}

	r, err := NewSiteRenderer()
	assertNoError(t, err)

	pages, err := r.BuildWebsite(data, 2)
	assertNoError(t, err)

	approvals.VerifyMap(t, pages)
}

func insertSlotData(phaseData *map[int]shared.PhaseData, inTier bool, loot shared.Loot) {
	entry := (*phaseData)[loot.Phase].Slots[loot.Slot]
	entry.InTier = true
	entry.Items = append(entry.Items, loot)
	(*phaseData)[loot.Phase].Slots[loot.Slot] = entry
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func createPhaseData() map[int]shared.PhaseData {

	data := map[int]shared.PhaseData{}

	for _, phase := range shared.PHASES {
		data[phase] = shared.PhaseData{}
		for _, slot := range shared.SLOTS {
			data[phase].Slots[int(slot)] = shared.SlotData{Items: []shared.Loot{}}
		}
	}

	return data
}
