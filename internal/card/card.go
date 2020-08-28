package card

import "sort"

// Card is defined by a value, a suit and a code which is an abreviation of the previous ones
type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// Values map including all posible ones for a card object
var Values = map[int]string{1: "ACE", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6",
	7: "7", 8: "8", 9: "9", 10: "0", 11: "JACK", 12: "QUEEN", 13: "KING"}

// Suits map including all posible ones for a card object
var Suits = map[int]string{1: "SPADES", 2: "DIAMONDS", 3: "CLUBS", 4: "HEARTS"}

// New creates a new card objet if the given parameters are correct
func New(value string, suit string) Card {
	c := Card{Value: value, Suit: suit, Code: string(value[0]) + string(suit[0])}
	return c
}

// Default creates a slice of sequential card objets
func Default() []Card {
	cards := []Card{}
	valuesKeys, SuitsKeys := Sort(Values), Sort(Suits)
	for _, s := range SuitsKeys {
		for _, k := range valuesKeys {
			cards = append(cards, New(Values[k], Suits[s]))
		}
	}
	return cards
}

// Sort given map keys and return them
func Sort(m map[int]string) []int {
	keys := make([]int, 0, len(m))
	for i := range m {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	return keys
}

// AllCodes creates a map of cards to identify each object by its code
func AllCodes() map[string]Card {
	cards := Default()
	cardCodes := map[string]Card{}
	for _, v := range cards {
		cardCodes[v.Code] = v
	}
	return cardCodes
}
