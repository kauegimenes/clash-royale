package lib

type Rarity struct {
	name     string // The name of the rarity
	cardsReq []int  // The number of cards needed for upgrading card at level "i"
	goldReq  []int  // The amount of gold needed for upgrading card at level "i"
	expGain  []int  // The amount of experience gained when upgrading card at level "i"
	goldCost []int  // The amount of gold needed to buy the "i + 1"-th card from the shop
}

func (r *Rarity) String() string {
	return r.name
}

func (r *Rarity) GetValue(attr Attribute) interface{} {
	switch attr {
	case CARDS_REQ:
		return r.cardsReq
	case GOLD_REQ:
		return r.goldReq
	case EXP_GAIN:
		return r.expGain
	default:
		return nil
	}
}

var (
	COMMON = &Rarity{
		name:     "Common",
		cardsReq: []int{1, 2, 4, 10, 20, 50, 100, 200, 400, 1000, 2000, 4000},
		goldReq:  []int{0, 5, 20, 50, 150, 400, 1000, 2000, 4000, 8000, 20000, 50000},
		expGain:  []int{0, 4, 5, 6, 10, 25, 50, 100, 200, 400, 800, 1600},
		goldCost: []int{
			3, 4, 5, 6, 7, 8, 10, 12, 14, 17,
			20, 24, 29, 35, 42, 50, 60, 72, 86, 103,
			124, 149, 179, 210, 258, 310, 372, 446, 535, 642,
			770, 924, 1109, 1331, 1597, 1916, 2299, 2759, 3311, 3973,
			4768, 5722, 6866, 8239, 9887, 11864, 14237, 17084, 20501, 24601,
		},
	}
	RARE = &Rarity{
		name:     "Rare",
		cardsReq: []int{1, 2, 4, 10, 20, 50, 100, 200, 400, 1000},
		goldReq:  []int{0, 50, 150, 400, 1000, 2000, 4000, 8000, 20000, 50000},
		expGain:  []int{0, 6, 10, 25, 50, 100, 200, 400, 800, 1600},
		goldCost: []int{
			40, 56, 78, 109, 153, 214, 300, 420, 588, 823,
			1152, 1613, 2258, 3161, 4425, 6195, 8673, 12142, 16999, 23799,
		},
	}
	EPIC = &Rarity{
		name:     "Epic",
		cardsReq: []int{1, 2, 4, 10, 20, 50, 100, 200},
		goldReq:  []int{0, 400, 1000, 2000, 4000, 8000, 20000, 50000},
		expGain:  []int{0, 25, 50, 100, 200, 400, 800, 1600},
		goldCost: []int{
			2000, 3600, 6480, 11664, 20995, 37791, 68024, 122443, 220397, 396715,
		},
	}
	LEGENDARY = &Rarity{
		name:     "Legendary",
		cardsReq: []int{1, 2, 4, 10, 20, 50},
		goldReq:  []int{0, 5000, 20000, 50000, 100000, 250000},
		expGain:  []int{0, 200, 400, 800, 1600, 3200},
		goldCost: []int{},
	}
)

var RARITIES = [...]*Rarity{
	COMMON,
	RARE,
	EPIC,
	LEGENDARY,
}
