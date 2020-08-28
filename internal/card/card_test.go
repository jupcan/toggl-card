package card

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	card := New("QUEEN", "HEARTS")
	assert.NotNil(t, card)
	assert.Equal(t, "QH", card.Code)
}

func TestDefault(t *testing.T) {
	var sig string
	cards := Default()
	for _, c := range cards {
		sig += c.Code
	}
	assert.Equal(t,
		"AS2S3S4S5S6S7S8S9S0SJSQSKSAD2D3D4D5D6D7D8D9D0DJDQDKDAC2C3C4C5C6C7C8C9C0CJCQCKCAH2H3H4H5H6H7H8H9H0HJHQHKH",
		sig)
}

func TestSort(t *testing.T) {
	values := map[int]string{3: "3", 12: "QUEEN", 6: "6", 1: "ACE", 9: "9"} 
	sortedKeys := Sort(values)
	assert.Equal(t, []int{1, 3, 6, 9, 12}, sortedKeys)
}

func TestAllCodes(t *testing.T) {
	cardCodes := AllCodes()
	assert.NotNil(t, cardCodes)
	assert.Equal(t, Card{Value:"KING", Suit:"CLUBS", Code:"KC"}, cardCodes["KC"])
	assert.Equal(t, "KING", cardCodes["KC"].Value)
	assert.Equal(t, "CLUBS", cardCodes["KC"].Suit)
}