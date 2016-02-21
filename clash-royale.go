package main

import (
	"fmt"
	"strings"
)

const X = -1

type Arena struct {
	id       int
	name     string
	trophies int
}

func (a *Arena) String() string {
	return fmt.Sprintf("Arena %d: %s", a.id, a.name)
}

var ARENAS = [...]Arena{
	Arena{id: 0, name: "Training Camp", trophies: 0},
	Arena{id: 1, name: "Goblin Stadium", trophies: 0},
	Arena{id: 2, name: "Bone Pit", trophies: 400},
	Arena{id: 3, name: "Barbarian Bowl", trophies: 800},
	Arena{id: 4, name: "P.E.K.K.A's Playhouse", trophies: 1100},
	Arena{id: 5, name: "Spell Valley", trophies: 1400},
	Arena{id: 6, name: "Royal Arena", trophies: 1700},
	Arena{id: 7, name: "Legendary Arena", trophies: 3000},
}

var (
	ARENA_0 = &ARENAS[0]
	ARENA_1 = &ARENAS[1]
	ARENA_2 = &ARENAS[2]
	ARENA_3 = &ARENAS[3]
	ARENA_4 = &ARENAS[4]
	ARENA_5 = &ARENAS[5]
	ARENA_6 = &ARENAS[6]
	ARENA_7 = &ARENAS[7]
)

//type RarityAttribute string
//
//const (
//	CARDS RarityAttribute = "Cards Required"
//	GOLD                  = "Gold Required"
//	EXP                   = "Experience Gained"
//)

type Rarity struct {
	id    int
	name  string
	cards []int
	gold  []int
	exp   []int
}

func (r *Rarity) String() string {
	return r.name
}

var RARITIES = [...]Rarity{
	Rarity{
		id:    0,
		name:  "Common",
		cards: []int{0, 2, 4, 10, 20, 50, 100, X, X, X, X, X},
		gold:  []int{0, 5, 20, 50, 150, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 4, 5, 6, 10, 25, 50, X, X, X, X, X},
	},
	Rarity{
		id:    1,
		name:  "Rare",
		cards: []int{0, 2, 4, 10, 20, 50, 100, X, X, X},
		gold:  []int{0, 50, 150, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 6, 10, 25, 50, X, X, X, X, X},
	},
	Rarity{
		id:    2,
		name:  "Epic",
		cards: []int{0, 2, 4, 10, 20, 50, 100, X},
		gold:  []int{0, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 6, 10, 25, 50, X, X, X},
	},
}

var (
	COMMON = &RARITIES[0]
	RARE   = &RARITIES[1]
	EPIC   = &RARITIES[2]
)

type Type string

const (
	TROOP    = Type("Troop")
	BUILDING = Type("Building")
	SPELL    = Type("Spell")
)

type Targets string

const (
	GROUND         = Targets("Ground")
	AIR_AND_GROUND = Targets("Air & Ground")
	BUILDINGS      = Targets("Buildings")
)

type Speed string

const (
	SLOW      = Speed("Slow")
	MEDIUM    = Speed("Medium")
	FAST      = Speed("Fast")
	VERY_FAST = Speed("Very Fast")
)

// Range:
const MELEE = 0

type CardAttribute struct {
	name      string
	isMutable bool
	format    func(x interface{}) string
}

func (attr *CardAttribute) String() string {
	return attr.name
}

var (
	NAME = &CardAttribute{
		"Name",
		false,
		func(x interface{}) string {
			return x.(string)
		},
	}
	ARENA = &CardAttribute{
		"Arena",
		false,
		func(x interface{}) string {
			return x.(*Arena).String()
		},
	}
	RARITY = &CardAttribute{
		"Rarity",
		false,
		func(x interface{}) string {
			return x.(*Rarity).String()
		},
	}
	TYPE = &CardAttribute{
		"Type",
		false,
		func(x interface{}) string {
			return string(x.(Type))
		},
	}
	DESC = &CardAttribute{
		"Description",
		false,
		func(x interface{}) string {
			return x.(string)
		},
	}
	COST = &CardAttribute{
		"Elixir Cost",
		false,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	HP = &CardAttribute{
		"Hitpoints",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	DPS = &CardAttribute{
		"Damage per Second",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	DAM = &CardAttribute{
		"Damage",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	ADAM = &CardAttribute{
		"Area Damage",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	DDAM = &CardAttribute{
		"Death Damage",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	SKE_LV = &CardAttribute{
		"Skeleton Level",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	SGO_LV = &CardAttribute{
		"Spear Goblin Level",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	HSPD = &CardAttribute{
		"Hit Speed",
		false,
		func(x interface{}) string {
			switch x.(type) {
			case int:
				return fmt.Sprintf("%dsec", x)
			case float64:
				return fmt.Sprintf("%.1fsec", x)
			default:
				return "???"
			}
		},
	}
	TGTS = &CardAttribute{
		"Targets",
		false,
		func(x interface{}) string {
			return string(x.(Targets))
		},
	}
	SPD = &CardAttribute{
		"Speed",
		false,
		func(x interface{}) string {
			return string(x.(Speed))
		},
	}
	RNG = &CardAttribute{
		"Range",
		false,
		func(x interface{}) string {
			switch x.(type) {
			case int:
				if x.(int) == MELEE {
					return "Melee"
				}
				return fmt.Sprintf("%d", x)
			case float64:
				return fmt.Sprintf("%.1f", x)
			default:
				return "???"
			}
		},
	}
	DTIME = &CardAttribute{
		"Deploy Time",
		false,
		func(x interface{}) string {
			switch x.(type) {
			case int:
				return fmt.Sprintf("%dsec", x)
			case float64:
				return fmt.Sprintf("%.1fsec", x)
			default:
				return "???"
			}
		},
	}
	COUNT = &CardAttribute{
		"Count",
		false,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
)

var CARD_ATTRIBUTES = [...]*CardAttribute{
	NAME,
	ARENA,
	RARITY,
	TYPE,
	DESC,
	COST,
	HP,
	DPS,
	DAM,
	ADAM,
	DDAM,
	SKE_LV,
	SGO_LV,
	HSPD,
	TGTS,
	SPD,
	RNG,
	DTIME,
	COUNT,
}

type Card map[*CardAttribute]interface{}

var CARDS_MAP = map[string]Card{
	"KNIGHT": Card{
		NAME:   "Knight",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "A tough melee fighter. The Barbarian's handsome, cultured cousin. Rumor has it that he was knighted based on the sheer awesomeness of his mustache alone.",
		COST:   3,
		HP:     []int{600, 660, 726, 798, 876, X, 1056, X, 1272, 1398, X, 1686},
		DPS:    []int{68, 74, 81, 90, 99, X, 120, X, 144, 158, X, 190},
		DAM:    []int{75, 82, 90, 99, 109, X, 132, X, 159, 174, X, 210},
		HSPD:   1.1,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    MELEE,
		DTIME:  1,
	},
	"ARCHERS": Card{
		NAME:   "Archers",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "A pair of unarmored ranged attackers. They'll help you with ground and air unit attacks, but you're on your own with coloring your hair.",
		COST:   3,
		HP:     []int{125, 137, 151, 166, X, 200, 220, X, 265, 291, 320, 351},
		DPS:    []int{33, 36, 40, 44, X, 53, 58, X, 70, 77, 85, 93},
		DAM:    []int{40, 44, 48, 53, X, 64, 70, X, 84, 93, 102, 112},
		HSPD:   1.2,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    5.5,
		DTIME:  1,
		COUNT:  2,
	},
	"BOMBER": Card{
		NAME:   "Bomber",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "Small, lightly protected skeleton that throws bombs. Deals damage that can wipe out a swarm of enemies.",
		COST:   3,
		HP:     []int{150, 165, 181, 199, 219, 240, 264, X, 318, 349, 384, 421},
		DPS:    []int{52, 57, 63, 70, 76, 84, 92, X, 111, 122, 134, 147},
		ADAM:   []int{100, 110, 121, 133, 146, 160, 176, X, 212, 233, 256, 281},
		HSPD:   1.9,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    5,
		DTIME:  1,
	},
	"GIANT": Card{
		NAME:   "Giant",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   "Slow but durable, only attacks buildings. A real one-man wrecking crew!",
		COST:   5,
		HP:     []int{2000, 2200, 2420, 2660, X, X, X, X, X, X},
		DPS:    []int{80, 88, 96, 106, X, X, X, 154, X, X},
		DAM:    []int{120, 132, 145, 159, X, X, X, 231, X, X},
		HSPD:   1.5,
		TGTS:   BUILDINGS,
		SPD:    SLOW,
		RNG:    MELEE,
		DTIME:  1,
	},
	"MUSKETEER": Card{
		NAME:   "Musketeer",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   "Don't be fooled by her delicately coiffed hair, the musketeer is a mean shot with her trusty boomstick.",
		COST:   4,
		HP:     []int{340, 374, 411, 452, 496, X, X, 656, X, X},
		DPS:    []int{81, 90, 98, 108, 119, X, X, X, X, X},
		DAM:    []int{90, 99, 108, 119, 131, X, X, X, X, X, X},
		HSPD:   1.1,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    6.5,
		DTIME:  1,
	},
}

var (
	KNIGHT    = CARDS_MAP["KNIGHT"]
	ARCHERS   = CARDS_MAP["ARCHERS"]
	BOMBER    = CARDS_MAP["BOMBER"]
	GIANT     = CARDS_MAP["GIANT"]
	MUSKETEER = CARDS_MAP["MUSKETEER"]
)

var CARDS = [...]Card{
	KNIGHT,
	ARCHERS,
	BOMBER,
	GIANT,
	MUSKETEER,
}

func main() {
	attrTitle := "Attribute"
	valueTitle := "Value"
	attrTitleLen := len(attrTitle)
	valueTitleLen := len(valueTitle)
	fixedAttrNameMaxLen := attrTitleLen
	mutableAttrNameMaxLen := attrTitleLen
	for _, attr := range CARD_ATTRIBUTES {
		attrNameLen := len(attr.name);
		if attr.isMutable {
			if attrNameLen > mutableAttrNameMaxLen {
				mutableAttrNameMaxLen = attrNameLen
			}
		} else {
			if attrNameLen > fixedAttrNameMaxLen {
				fixedAttrNameMaxLen = attrNameLen
			}
		}
	}

	for _, card := range CARDS {
		fmt.Printf("%*s | %s\n", -fixedAttrNameMaxLen, attrTitle, valueTitle)
		fmt.Printf("%s | %s\n", strings.Repeat("-", fixedAttrNameMaxLen), strings.Repeat("-", valueTitleLen))
		for _, attr := range CARD_ATTRIBUTES {
			if attr.isMutable {
				continue
			}
			if value, ok := card[attr]; ok {
				fmt.Printf("%*s | %s\n", -fixedAttrNameMaxLen, attr, attr.format(value))
			}
		}
		fmt.Println()
	}
}