package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten", "jack", "queen", "king":
		value = 10
	default:
		// do nothing
	}

	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	action := "H"
	switch {
	case playerValue == 22:
		action = "P"
	case playerValue == 21 && dealerValue < 10:
		action = "W"
	case playerValue == 21:
		action = "S"
	case playerValue >= 17 && playerValue <= 20:
		action = "S"
	case playerValue >= 12 && playerValue <= 16 && dealerValue < 7:
		action = "S"
	case playerValue >= 12 && playerValue <= 16:
		action = "H"
	case playerValue <= 11:
		action = "H"
	default:
		// do nothing
	}

	return action
}
