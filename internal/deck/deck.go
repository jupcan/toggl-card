package deck

import (
	"errors"
	"math/rand"
	"github.com/google/uuid"
	"toggl-card/internal/card"
)

// Deck is defined as a cobination of card objects
type Deck struct {
	ID        uuid.UUID   `json:"deck_id"`
	Shuffled  bool        `json:"shuffled"`
	Remaining int         `json:"remaining"`
	Cards     []card.Card `json:"cards"`
}

// Decks is a slice to store decks since no persistence is implemented
var Decks []Deck

// New creates a default sequential card objets deck
func New(shuffle bool) Deck {
	cards := card.Default()
	return NewPartial(shuffle, cards)
}

// NewPartial creates a custom card objets deck
func NewPartial(shuffle bool, cards []card.Card) Deck {
	return Deck{ID: uuid.New(), Shuffled: shuffle, Remaining: len(cards), Cards: cards}
}

// SetCards of a deck 
func (d *Deck) SetCards(c []card.Card) {
	d.Cards = c
}

// Signature to represent a deck and its cards order 
func (d *Deck) Signature() (string) {
	var sig string
	for _, card := range d.Cards {
		sig += card.Code
	}
	return sig
}

// Shuffle deck's cards following Fisher-Yates algorithm
func (d *Deck) Shuffle() {
	cards := d.Cards
	for i := len(cards) - 1; i > 0; i-- {
		r := rand.Intn(i + 1)
		cards[r], cards[i] = cards[i], cards[r]
	}
}

// Draw n cards from the deck if enough cards in it
func (d *Deck) Draw(n int) ([]card.Card, error) {
	if len(d.Cards) < n {
		return []card.Card{}, errors.New("Error. Not enough cards in the deck")
	}
	cards := d.Cards
	drawn := cards[:n]
	d.SetCards(cards[n:])
	return drawn, nil
}