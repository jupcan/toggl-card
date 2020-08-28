package server

import (
	"testing"
	"net/http"
	"encoding/json"
	"net/http/httptest"
	"toggl-card/internal/deck"
	"github.com/stretchr/testify/assert"
)

func Request(method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	server := New()
	server.Router().ServeHTTP(w, req)
	return w
}

func TestCreate(t *testing.T) {
	t.Parallel()
	w := Request("POST", "/deck/create")
	var deck deck.Deck
	err := json.Unmarshal([]byte(w.Body.String()), &deck)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 52, deck.Remaining)
	assert.Equal(t, false, deck.Shuffled)
}

func TestOpen(t *testing.T) {
	t.Parallel()
	w := Request("POST", "/deck/create")
	assert.Equal(t, http.StatusOK, w.Code)

	var deck deck.Deck
	err := json.Unmarshal([]byte(w.Body.String()), &deck)
	w = Request("GET", "/deck?uuid=" + deck.ID.String())

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDraw(t *testing.T) {
	t.Parallel()
	w := Request("POST", "/deck/create")
	assert.Equal(t, http.StatusOK, w.Code)

	var deck deck.Deck
	err := json.Unmarshal([]byte(w.Body.String()), &deck)

	// Expected errors if no deck is found, not passed over or not enough cards left
	w = Request("GET", "/deck/" + deck.ID.String() + "/draw?count=60")
	assert.Equal(t, http.StatusBadRequest, w.Code)
	w = Request("GET", "/deck/draw?count=5")
	assert.Equal(t, http.StatusNotFound, w.Code)
	w = Request("GET", "/deck/123/draw?count=5")
	assert.Equal(t, http.StatusNotFound, w.Code)

	w = Request("GET", "/deck/" + deck.ID.String() + "/draw?count=2")
	assert.Equal(t, http.StatusOK, w.Code)
	
	w = Request("GET", "/deck?uuid=" + deck.ID.String())
	assert.Equal(t, http.StatusOK, w.Code)
	err = json.Unmarshal([]byte(w.Body.String()), &deck)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 50, deck.Remaining)
}