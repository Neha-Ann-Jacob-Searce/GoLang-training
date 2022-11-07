package main

//create a new type of deck which is a slice of strings
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck { // receiver not needed
	cards := deck{}
	cardSuits := []string{"spades", "diamonds", "cloves", "clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}
	for _, suit := range cardSuits {
		for value := range cardValue {
			cards = append(cards, cardValue[value]+" of "+suit)
		}
	}
	return cards
}

// reveiver sets up moethods on variables that we create
// receiver function
func (d deck) print() { // The receiver here is of type deck. Any variable inthe application of type deck now gets access to the method print(); the variable d before deck is an instance of the deck variable. Its a reference to the deck variacle
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // multiple return
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("\nERROR: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() { // random number generation
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		newPos := r.Intn(len(d) - 1)
		d[index], d[newPos] = d[newPos], d[index]
	}
}
