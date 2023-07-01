package tmbjsonparse

import (
	"encoding/json"
	"time"

	"github.com/jaskian/tmb-tier-site/shared"
)

func unmarshalTMBJson(jsonData []byte) (tmbData, error) {
	tmbdata := tmbData{}
	err := json.Unmarshal(jsonData, &tmbdata)

	if err != nil {
		return tmbdata, err
	}

	return tmbdata, nil
}

func (primaryExport *tmbData) mergeTMBData(oldPhases map[int]tmbData) {
	for _, data := range oldPhases {
		for _, oChar := range data {
			for i, pChar := range *primaryExport {
				if pChar.Name == oChar.Name {
					(*primaryExport)[i].Wishlisted = append((*primaryExport)[i].Wishlisted, oChar.Wishlisted...)
				}
			}
		}
	}
}

type tmbData []character

type character struct {
	Name         string `json:"name"`
	Class        string `json:"class"`
	Spec         string `json:"spec"`
	ReceivedLoot []loot `json:"received"`
	Wishlisted   []loot `json:"wishlist"`
}

type loot struct {
	ItemID        int    `json:"item_id"`
	ItemName      string `json:"name"`
	InventoryType int    `json:"inventory_type"`
	InstanceID    int    `json:"instance_id"`
	Pivot         pivot  `json:"pivot"`
}

type pivot struct {
	Date           string `json:"received_at"`
	Offspec        int    `json:"is_offspec"`
	OfficerNote    string `json:"officer_note"`
	ReceivedWLItem int    `json:"is_received"`
}

func NewLoot(i loot, phase int, slot int) shared.Loot {
	t := time.Unix(0, 0)

	if i.Pivot.Date != "" {
		t, _ = time.Parse(TMB_TIME_FORMAT, i.Pivot.Date)
	}

	result := shared.Loot{
		ItemID:   i.ItemID,
		ItemName: i.ItemName,
		Phase:    phase,
		Slot:     slot,
		Date:     t,
	}

	return result
}
