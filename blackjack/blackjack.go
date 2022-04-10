package blackjack

import "fmt"

var Deck = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"other": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return Deck[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1 := Deck[card1]
	c2 := Deck[card2]
	cardDealer := Deck[dealerCard]

	if c1 == c2 && c1 == 11 {
		return "P"
	}

	blackjack := int(c1+c2) == 21
	rangeCards := int(c1 + c2)
	autoWin := cardDealer != 11 && cardDealer != 10
	if blackjack {
		if autoWin {
			return "W"
		} else {
			return "S"
		}
	}
	if rangeCards >= 17 && rangeCards <= 20 {
		return "S"
	}

	if rangeCards >= 12 && rangeCards <= 16 && cardDealer >= 7 {
		return "H"
	}

	if rangeCards >= 12 && rangeCards <= 16 && cardDealer < 7 {
		return "S"
	}

	if rangeCards <= 11 {
		return "H"
	}

	return "H"

	// Stand (S)
	// Hit (H)
	// Split (P)
	// Automatically win (W)

}

func main() {
	// value := ParseCard("nine")
	// fmt.Println(value)
	// Output: 11
	//fmt.Println(FirstTurn("ace", "ace", "jack"))
	//Output: "P"
	//fmt.Println(FirstTurn("ace", "king", "ace"))
	//Output: "S"
	fmt.Println(FirstTurn("five", "queen", "ace"))
	//Output: "H"
}
