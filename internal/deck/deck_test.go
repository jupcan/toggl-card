package deck

import (
	"testing"
	"toggl-card/internal/card"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	deck := New(false)
	assert.NotNil(t, deck)
	assert.Equal(t, 52, len(deck.Cards))
}

func TestNewPartial(t *testing.T) {
	cards := []card.Card{
		card.New("ACE", "SPADES"),
		card.New("KING", "DIAMONDS"),
		card.New("ACE", "CLUBS"),
		card.New("2", "CLUBS"),
		card.New("KING", "HEARTS"),
	}
	deck := NewPartial(false, cards)
	assert.Equal(t, 5, len(deck.Cards))
	for i := range deck.Cards{
		assert.Equal(t, deck.Cards[i].Code, cards[i].Code)
	}
}

func TestSignature(t *testing.T) {
	cardCodes := card.AllCodes()
	deck := NewPartial(false, []card.Card{cardCodes["AS"], cardCodes["KD"], cardCodes["AC"]})
	assert.Equal(t, "ASKDAC", deck.Signature())
}

func TestShuffle(t *testing.T) {
	unshuffled, shuffled  := New(false), New(false)
	assert.Equal(t, unshuffled.Signature(), shuffled.Signature())
	shuffled.Shuffle()
	assert.NotEqual(t, unshuffled.Signature(), shuffled.Signature())
}

func TestDraw(t *testing.T) {
	deck := New(true)
	for i := 0; i < 52; i++ {
		drawn, err := deck.Draw(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(drawn))
	}
	assert.Equal(t, 0, len(deck.Cards))
}