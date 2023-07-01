package tmbjsonparse

import (
	"embed"
	"reflect"
	"testing"
)

var (
	//go:embed data/slim.json
	example      []byte
	importKeeper embed.FS
)

func TestUnmarshalTMBJson(t *testing.T) {

	want := character{
		Name:  "Asara",
		Class: "Druid",
		Spec:  "Feral",
		ReceivedLoot: []loot{
			{
				ItemID:        40208,
				ItemName:      "Cryptfiend's Bite",
				InventoryType: 17,
				InstanceID:    20,
				Pivot: pivot{
					Date: "2022-11-10 00:00:00",
				},
			},
			{
				ItemID:        40250,
				ItemName:      "Aged Winter Cloak",
				InventoryType: 16,
				InstanceID:    20,
				Pivot: pivot{
					Date:    "2022-11-03 00:00:00",
					Offspec: 1,
				},
			},
		},
		Wishlisted: []loot{
			{
				ItemID:        45224,
				ItemName:      "Drape of the Lithe",
				InventoryType: 16,
				InstanceID:    20,
				Pivot: pivot{
					ReceivedWLItem: 1,
				},
			},
		},
	}

	got, err := unmarshalTMBJson(example)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got[0], want) {
		t.Errorf("got %v, want %v", got[0], want)
	}
}

func TestMergeTMBData(t *testing.T) {

	primary := &tmbData{
		{
			Name: "Asara",
			Wishlisted: []loot{
				{
					ItemID:        70609,
					ItemName:      "Reign of the Unliving",
					InventoryType: 12,
					InstanceID:    32,
					Pivot: pivot{
						ReceivedWLItem: 0,
					},
				},
			},
		},
	}

	oldPhaseData := map[int]tmbData{
		1: tmbData{
			{
				Name: "Asara",
				Wishlisted: []loot{
					{
						ItemID:        40383,
						ItemName:      "Calamity's Grasp",
						InventoryType: 21,
						InstanceID:    20,
						Pivot: pivot{
							ReceivedWLItem: 0,
						},
					},
				},
			},
		},
		2: tmbData{
			{
				Name: "Asara",
				Wishlisted: []loot{
					{
						ItemID:        45224,
						ItemName:      "Drape of the Lithe",
						InventoryType: 16,
						InstanceID:    28,
						Pivot: pivot{
							ReceivedWLItem: 1,
						},
					},
				},
			},
		},
	}

	primary.mergeTMBData(oldPhaseData)

	if len((*primary)[0].Wishlisted) != 3 {
		t.Errorf("Expected 3, got %d", len((*primary)[0].Wishlisted))
	}

}
