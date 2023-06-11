package shared

type Slot int

const (
	Head Slot = iota + 1
	Neck
	Shoulder
	Shirt
	Chest
	Belt
	Legs
	Feet
	Wrist
	Gloves
	Ring
	Trinket
	Cloak     Slot = 16
	TwoHander Slot = 17
	OneHander Slot = 18
	Offhand   Slot = 19
	Ranged    Slot = 20
)

var SLOTS = []Slot{
	Head,
	Neck,
	Shoulder,
	Shirt,
	Chest,
	Belt,
	Legs,
	Feet,
	Wrist,
	Gloves,
	Ring,
	Trinket,
	Cloak,
	TwoHander,
	OneHander,
	Offhand,
	Ranged,
}

var InventoryTypeMappings = map[int]Slot{

	20: Chest, // chests? e.g. Sympathy

	13: OneHander, // 1H e.g. Malice
	21: OneHander, // MH e.g. Golden Saronite Dragon

	14: Offhand, // Shield e.g. Wall of Terror
	22: Offhand, // OH e.g. Delirium's Touch
	23: Offhand, // OH non-wep e.g. Cosmos

	15: Ranged, // Bow e.g. Arrowsong
	25: Ranged, // Thrown e.g. Spinning Fate
	26: Ranged, // Ranged e.g. Final Voyage
	28: Ranged, // Libram etc e.g. Libram of Resurgence

	17: TwoHander, // 2H e.g. Worldcarver
}
