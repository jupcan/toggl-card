package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"toggl-card/pkg/server"
)

func init() {
	// Seed rand with current time when using it to shuffle a deck
	rand.Seed(time.Now().UnixNano())
}

func main() {
	s := server.New()
	fmt.Println("Toggl Backend Unattended Programming Test REST API")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}