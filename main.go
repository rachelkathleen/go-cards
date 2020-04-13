package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i+1, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
