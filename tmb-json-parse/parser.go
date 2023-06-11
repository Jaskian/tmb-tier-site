package tmbjsonparse

import (
	"encoding/json"
)

func unmarshalTMBJson(jsonData []byte) (tmbData, error) {
	tmbdata := tmbData{}
	err := json.Unmarshal(jsonData, &tmbdata)

	if err != nil {
		return tmbdata, err
	}

	return tmbdata, nil
}

type tmbData []character

type character struct {
	Name         string `json:"name"`
	Class        string `json:"class"`
	Spec         string `json:"spec"`
	ReceivedLoot []loot `json:"received"`
}

type loot struct {
	ItemID        int    `json:"item_id"`
	ItemName      string `json:"name"`
	InventoryType int    `json:"inventory_type"`
	InstanceID    int    `json:"instance_id"`
	Pivot         pivot  `json:"pivot"`
}

type pivot struct {
	Date string `json:"received_at"`
}
