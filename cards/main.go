package main

import "fmt"

func main() {

	if false {
		study()
	}

	cards := newDeck()
	// for _, card := range cards {
	// 	fmt.Println(card)
	// }

	cards.print()
	cards.shuffle()
	cards.print()

}

func study() {
	cards := newDeck()
	for i, card := range cards {
		fmt.Println(i, card)
	}

	d1, d2 := deal(cards, 5)

	for _, card := range d1 {
		fmt.Println(card)
	}
	fmt.Println("==================")
	for _, card := range d2 {
		fmt.Println(card)
	}

	g := "hi here"
	fmt.Println([]byte(g))

	dd := newDeck()
	dd.saveToFile("hello")
	// fmt.Println([]byte(dd))

}

// gopath 설정 해주기
