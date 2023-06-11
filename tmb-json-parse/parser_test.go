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
					Date: "2022-11-03 00:00:00",
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
	// var sortedLoot = map[int][]Loot{}

	// for _, c := range tmbdata {
	// 	for _, l := range c.ReceivedLoot {
	// 		sortedLoot[l.InventoryType] = append(sortedLoot[l.InventoryType], l)
	// 	}
	// }

	// for i, l := range sortedLoot {
	// 	fmt.Printf("=== Slot %d ===", i)
	// 	for _, li := range l {
	// 		fmt.Println(li.ItemName)
	// 	}
	// }

	// var sortedLootInstance = map[int][]loot{}

	// for _, c := range tmbdata {
	// 	for _, l := range c.ReceivedLoot {
	// 		sortedLootInstance[l.InstanceID] = append(sortedLootInstance[l.InstanceID], l)
	// 	}
	// }

	// for i, l := range sortedLootInstance {
	// 	fmt.Printf("=== Instance %d ===", i)
	// 	for _, li := range l {
	// 		fmt.Println(li.ItemName)
	// 	}
	// }
}
