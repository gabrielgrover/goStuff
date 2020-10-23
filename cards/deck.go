package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	data := []byte(d.toString())

	return ioutil.WriteFile(fileName, data, 0666)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	data := string(bs)

	return strings.Split(data, ",")
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
