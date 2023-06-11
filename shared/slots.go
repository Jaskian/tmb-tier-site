package shared

type Slot int

const (
	Head Slot = iota + 1
	Neck
	Shoulder
	//Shirt
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
	Head:     "https://static.wikia.nocookie.net/wowpedia/images/c/c3/Ui-paperdoll-slot-head.png/revision/latest?cb=20070607015044",
	Neck:     "https://static.wikia.nocookie.net/wowpedia/images/d/d4/Ui-paperdoll-slot-neck.png/revision/latest?cb=20070607015138",
	Shoulder: "https://static.wikia.nocookie.net/wowpedia/images/f/fa/Ui-paperdoll-slot-shoulder.png/revision/latest?cb=20070607015347",
	//Shirt : "",
	Chest:     "https://static.wikia.nocookie.net/wowpedia/images/b/b7/Ui-paperdoll-slot-chest.png/revision/latest?cb=20070606225854",
	Belt:      "https://static.wikia.nocookie.net/wowpedia/images/c/cd/Ui-paperdoll-slot-waist.png/revision/latest?cb=20070607015409",
	Legs:      "https://static.wikia.nocookie.net/wowpedia/images/1/14/Ui-paperdoll-slot-legs.png/revision/latest?cb=20070607015103",
	Feet:      "https://static.wikia.nocookie.net/wowpedia/images/a/a5/Ui-paperdoll-slot-feet.png/revision/latest?cb=20070607015023",
	Wrist:     "https://static.wikia.nocookie.net/wowpedia/images/1/1d/Ui-paperdoll-slot-wrists.png/revision/latest?cb=20070607015415",
	Gloves:    "https://static.wikia.nocookie.net/wowpedia/images/2/22/Ui-paperdoll-slot-hands.png/revision/latest?cb=20070607015036",
	Ring:      "https://static.wikia.nocookie.net/wowpedia/images/c/c2/Ui-paperdoll-slot-finger.png/revision/latest?cb=20070607015031",
	Trinket:   "https://static.wikia.nocookie.net/wowpedia/images/2/26/Ui-paperdoll-slot-trinket.png/revision/latest?cb=20070607015403",
	Cloak:     "https://wow.zamimg.com/images/wow/icons/large/inv_misc_cape_03.jpg",
	TwoHander: "https://static.wikia.nocookie.net/wowpedia/images/f/f5/Ui-paperdoll-slot-mainhand.png/revision/latest?cb=20070607015117",
	OneHander: "https://static.wikia.nocookie.net/wowpedia/images/f/f5/Ui-paperdoll-slot-mainhand.png/revision/latest?cb=20070607015117",
	Offhand:   "https://static.wikia.nocookie.net/wowpedia/images/3/30/Ui-paperdoll-slot-secondaryhand.png/revision/latest?cb=20070607015335",
	Ranged:    "https://static.wikia.nocookie.net/wowpedia/images/9/9f/Ui-paperdoll-slot-ranged.png/revision/latest?cb=20070607015146",
}

var SLOTS = []Slot{
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
