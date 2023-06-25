package shared

import "time"

type TMBData []Character

type Character struct {
	Name     string
	Class    string
	Spec     string
	Phases   map[int]PhaseData
	Wishlist Wishlist
	KeyItems KeyItems
}

type PhaseData map[int]SlotData

type SlotData struct {
	InTier bool
	Items  []Loot
}

type Loot struct {
	ItemID   int
	ItemName string
	Phase    int
	Slot     int
	Date     time.Time
}
type Wishlist struct {
	Received     int
	Total        int
	WishlistLoot []WishlistLoot
}
type WishlistLoot struct {
	Loot
	Received bool
}

type KeyItems struct {
	Trophies int
}
