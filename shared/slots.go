package shared

type Slot int

const (
	Misc      = 0
	Head Slot = iota
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

var SLOT_IMAGE_URLS = map[Slot]string{
	Misc:     "https://wow.zamimg.com/images/wow/icons/large/inv_misc_questionmark.jpg",
	Head:     "https://wow.zamimg.com/images/wow/icons/large/inv_helmet_37.jpg",
	Neck:     "https://wow.zamimg.com/images/wow/icons/large/inv_jewelry_necklace_15.jpg",
	Shoulder: "https://wow.zamimg.com/images/wow/icons/large/inv_shoulder_14.jpg",
	//Shirt : "",
	Chest:     "https://wow.zamimg.com/images/wow/icons/large/inv_chest_chain_15.jpg",
	Belt:      "https://wow.zamimg.com/images/wow/icons/large/inv_belt_15.jpg",
	Legs:      "https://wow.zamimg.com/images/wow/icons/large/inv_pants_03.jpg",
	Feet:      "https://wow.zamimg.com/images/wow/icons/large/inv_boots_07.jpg",
	Wrist:     "https://wow.zamimg.com/images/wow/icons/large/inv_bracer_05.jpg",
	Gloves:    "https://wow.zamimg.com/images/wow/icons/large/inv_gauntlets_27.jpg",
	Ring:      "https://wow.zamimg.com/images/wow/icons/large/inv_jewelry_ring_03.jpg",
	Trinket:   "https://wow.zamimg.com/images/wow/icons/large/inv_jewelry_talisman_10.jpg",
	Cloak:     "https://wow.zamimg.com/images/wow/icons/large/inv_misc_cape_05.jpg",
	TwoHander: "https://wow.zamimg.com/images/wow/icons/large/inv_sword_50.jpg",
	OneHander: "https://wow.zamimg.com/images/wow/icons/large/inv_sword_39.jpg",
	Offhand:   "https://wow.zamimg.com/images/wow/icons/large/inv_misc_book_07.jpg",
	Ranged:    "https://wow.zamimg.com/images/wow/icons/large/inv_weapon_bow_01.jpg",
}

var SLOTS = []Slot{
	Misc,
	Head,
	Neck,
	Shoulder,
	//Shirt,
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

var TokenMapping = map[string]Slot{
	"Crown":       Head,
	"Helm":        Head,
	"Head":        Head,
	"Shoulder":    Shoulder,
	"Shoulders":   Shoulder,
	"Mantle":      Shoulder,
	"Spaulders":   Shoulder,
	"Chest":       Chest,
	"Chestguard":  Chest,
	"Regalia":     Chest,
	"Breastplate": Chest,
	"Hands":       Gloves,
	"Gauntlets":   Gloves,
	"Gloves":      Gloves,
	"Legs":        Legs,
	"Legplates":   Legs,
	"Leggings":    Legs,
}
