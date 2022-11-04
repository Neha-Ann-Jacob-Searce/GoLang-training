package main

import "fmt"

func main() {
	var hello string = "Hello there!"
	fmt.Println([]byte(hello))
	//card := newCard() // only while initialisation := is to notify that go should identify datatype of the variable
	//card = "Five of Diamonds"
	cards := newDeck()
	cards.print()
	//fmt.Println(card)

	// cards := deck{"Ace of diamonds", newCard()} // slice of type string
	// fmt.Println(cards)
	// cards = append(cards, "three of hearts") // append does nt modify the existing slice. Instead, it returns a new slice. The new slice is assigned back to the variable of cards
	// fmt.Println(cards)

	//iteration over a slice
	fmt.Println("Type 1")
	for index, card := range cards {
		fmt.Println(index, card)
	}

	fmt.Println("Type 2")
	for card := range cards {
		fmt.Println(cards[card])
	}

	fmt.Println("Type 3")
	for card := range cards {
		fmt.Println(card)
	}

	fmt.Println("Way 2")
	cards.print()

	// splitting cards slice
	fmt.Println("\nHand")
	hand, remainingDeck := deal(cards, 5)
	fmt.Println(hand.toString())
	fmt.Println([]byte(hand.toString()))
	fmt.Println("\nRemaining deck")
	remainingDeck.print()

	cards.saveToFile("my_cards")
	fmt.Println()
	newCards := newDeckFromFile("my_cards")
	newCards.print()

	fmt.Println("\nShuffling")
	newCards.shuffle()
	newCards.print()

	fmt.Println("\n\nERROR HANDLING")
	newCards = newDeckFromFile("my")
	newCards.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
