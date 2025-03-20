package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// ______________________________________________________________________
// Create a new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// ______________________________________________________________________
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// ______________________________________________________________________
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// ______________________________________________________________________
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// ______________________________________________________________________
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// _____________________________________________________________________
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// println(string(bs))
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// _____________________________________________________________________
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for idx := range d {
		newPosition := r.Intn(len(d) - 1)
		d[idx], d[newPosition] = d[newPosition], d[idx]
	}
}
