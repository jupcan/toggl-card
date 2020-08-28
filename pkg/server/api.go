package server

import (
	"strconv"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"toggl-card/internal/card"
	"toggl-card/internal/deck"
)

type api struct {
	router http.Handler
}

// Server is the interface that wraps the router method
type Server interface {
	Router() http.Handler
}

// New implements Server iface and creates a router to handle api endpoints
func New() Server {
	a := &api{}
	r := mux.NewRouter()

	// Endpoints to be handled by the router
	r.HandleFunc("/deck", a.Open).Methods("GET")
	r.HandleFunc("/deck/create", a.Create).Methods("POST")
	r.HandleFunc("/deck/{uuid}/draw", a.Draw).Methods("GET")
	a.router = r
	return a
}

// Router returns the router of a given api
func (a *api) Router() http.Handler {
	return a.router
}

// Create a deck of cards accepting shuffle and custom cards parameters
func (a *api) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	shuffle := r.URL.Query().Get("shuffle")
	cardsStr := r.URL.Query().Get("cards")
	var d deck.Deck

	if len(cardsStr) != 0 {
		cardsLst := strings.Split(cardsStr, ",")
		var cards []card.Card
		cardCodes := card.AllCodes()
		for _, code := range cardsLst {
			if v, ok := cardCodes[code]; ok {
				cards = append(cards, v)
			}
		}
		d = deck.NewPartial(false, cards)	
	} else {
		if shuffle == "true" {
			d = deck.New(true)
			d.Shuffle()
		} else {
			d = deck.New(false)
		}
	}
	
	deck.Decks = append(deck.Decks, d)
	json.NewEncoder(w).Encode(struct {
		ID        uuid.UUID   `json:"deck_id"`
		Shuffled  bool        `json:"shuffled"`
		Remaining int         `json:"remaining"`
	}{d.ID, d.Shuffled, d.Remaining})
}

// Open a given deck by ist UUID and return all its properties 
func (a *api) Open(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uuid := r.URL.Query().Get("uuid")
	if len(uuid) != 0 {
		for _, deck := range deck.Decks {
			if deck.ID.String() == uuid {
				json.NewEncoder(w).Encode(deck)
				return
			}
		}
		http.Error(w, "404. Deck with provided UUID could not be found.", 404)
	} else {
		http.Error(w, "400. Missing UUID parameter.", 400)
	}
}

// Draw n cards from a Deck that must exist or be passed over
func (a *api) Draw(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	count := r.URL.Query().Get("count")
	n, countErr := strconv.Atoi(count)
	if countErr == nil && n > 0 {
		for i, d := range deck.Decks {
			if d.ID.String() == uuid {
				drawn, drawErr := d.Draw(n)
				// Update deck from Decks slice with remaining cards 
				deck.Decks[i].Cards = d.Cards
				deck.Decks[i].Remaining = len(d.Cards)

				if drawErr == nil {
					json.NewEncoder(w).Encode(struct {
						Cards []card.Card `json:"cards"`
					}{drawn})
					return
				} 
				http.Error(w, "400. Not enough cards in the deck", 400)
				return			
			}
		}
		http.Error(w, "404. Deck with provided UUID could not be found.", 404)
	} else {
		http.Error(w, "400. Count parameter with number of cards to be drawn needs to be provided.", 400)
	}
}