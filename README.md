# [Toggl-card](/doc/Toggl_Backend_Unattended_Programming_Test.pdf)
[![uclm](https://img.shields.io/badge/personal-project-red.svg?&longCache=true&colorA=27a79a&colorB=555555&style=for-the-badge)](http://www.juanperea.me)  
> REST API to handle decks and cards to be used in any game like Poker and Blackjack.

![toggl-banner](doc/banner.jpg)

## Installation
Assigment developed using **Go**. A simple makefile is provided to ease, as its help parameter states, some tasks:

```
all      run all tests, builds and executes the api
build    builds api binary file
test     run all implemented tests
clean    cleans and deletes current binary
run      builds binary file and runs it
get      get all dependencies and install them 
doc      generate documentation from code commets
```

Dependencies are listed in the go module and can be installed running **make get**. Mainly 3 of them have been used:

1. [github.com/google/uuid](https://github.com/google/uuid) - to create UUIDs that unequivocally represent decks
2. [github.com/gorilla/mux](https://github.com/google/uuid) - to implement a router and dispatcher for matching incoming requests to their handler
3. [github.com/stretchr/testify](https://github.com/google/uuid) - to make easier the tests development by the use of common assertions functions

Use **make all** to run all tests, see the results, build de api binary file and execute it. The router will be listening for requests in localhost 8080 port by default.

## Usage

The are 3 different functionalities implemented:

**1. Create Deck (POST)** -  creates a new standard 52-card French playing cards deck. It includes all thirteen ranks in each of the four suits: clubs :clubs:, diamonds :diamonds:, hearts :hearts: and spades :spades: and allows the following options to the request:

- to be *shuffled* or not — by default the deck is sequential: A spades, 2-spades, 3-spades... followed by diamonds, clubs, then hearts. If <span style="background-color: #D9D9D9">?shuffle=true</span> is provided, the deck created will be randomly shuffled in execution time.
- to be *full* or *partial* — by default as stated above, it returns the standard 52 cards, otherwise the request would accept the wanted cards like this example: <span style="background-color: #D9D9D9">?cards=AS,KD,AC,2C,KH</span> (ACE of SPADES, KING of DIAMONDS...)

```
/deck/create
```
```json
{
    "deck_id": "830a6100-766a-405f-be6b-1d3b63b4c9d2",
    "shuffled": false,
    "remaining": 52
}
```
```
/deck/create?shuffle=<true/false>
/deck/create?shuffle=true
```
```json
{
    "deck_id": "e320a37e-9f91-4fbb-91bc-995dc8534346",
    "shuffled": true,
    "remaining": 52
}
```
```
/deck/create?cards=<CARCODE,CARCODE...>
/deck/create?cards=AS,KD,AC,2C,KH
```
```json
{
    "deck_id": "158e7d53-03fa-4677-9e08-6aef7c299266",
    "shuffled": false,
    "remaining": 5
}
```

**2. Open Deck (GET)** - returns a given deck by its UUID, which must be passed over and returning an error if the deck does not exists. The method opens the deck by listing all its properties and cards. Example of request and response:

```
/deck?uuid=<uuid>
/deck?uuid=b0c4147e-17dd-40d5-b3d2-4714fc20434a
```
```json
{
    "deck_id": "b0c4147e-17dd-40d5-b3d2-4714fc20434a",
    "shuffled": false,
    "remaining": 5,
    "cards": [
        {
            "value": "KING",
            "suit": "HEARTS",
            "code": "KH"
        },
        {
            "value": "7",
            "suit": "SPADES",
            "code": "7S"
        },
        {
            "value": "QUEEN",
            "suit": "DIAMONDS",
            "code": "QD"
        },
    ]
}
```
**3. Draw Card (GET)** - draws a card(s) of a given Deck. If the deck is not passed over or invalid it returns an error. A count parameter needs to be provided to define how many cards to draw and it should be greater than 0. 

```
/deck/<uuid>/draw?count=<int>
/deck/b0c4147e-17dd-40d5-b3d2-4714fc20434a/draw?count=2
```
```json
{
    "cards": [
        {
            "value": "KING",
            "suit": "CLUBS",
            "code": "KC"
        },
        {
            "value": "5",
            "suit": "SPADES",
            "code": "5S"
        }
    ]
}
```

## Information

Some considerations regarding the development of the REST API.

- Despite not beeing a big project nor having too much files, have decided to use a proper directory structure to organise all the source code and tests as if was going to grow in the future. Same reasoning when thinking of the creation of **deck** and **card** packages, not a must and could have implemented all in a main one but thought it is more organised this way. For example object creation calls are more idiomatic — *deck.New()* or *card.New()* — and follow [effective Go](https://golang.org/doc/effective_go.html) rather than *NewCard()* or *NewDeck()*
- Have implemented Fisher–Yates shuffle algorithm to randomize a deck in O(n) time
- In main init function, I have set rand seed to do it only once and not in every shuffle iteration
- I decided to use **maps** as the data strcuture for available card values and suits so that its sorting is easier than if using a slice of strings. The **cardCodes** map representing each possible card by its code could also help to qucikly find a player with a specific card, something used to start some card games
- To do not work outside the required scope of the assigment have not implemented persistence, thus a slice is used to store all created decks so when the execution stops all of them are lost 
- 0 is the value used for each suit 10th card to represent it in a two-digit code as the other ones and not to be confused with the ace, usually known as 1
