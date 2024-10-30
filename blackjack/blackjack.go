package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	case "other":
		return 0
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	v1 := ParseCard(card1)
	v2 := ParseCard(card2)
	v3 := ParseCard(dealerCard)

	if v1 == 11 && v2 == 11 {
		return "P"
	}

	if v1+v2 == 21 {
		if v3 == 11 || v3 == 10 {
			return "S"
		} else {
			return "W"
		}
	}

	if v1+v2 >= 17 && v1+v2 <= 20 {
		return "S"
	}

	if v1+v2 >= 12 && v1+v2 <= 16 {
		if v3 >= 7 {
			return "H"
		}
		return "S"
	}

	if v1+v2 <= 11 {
		return "H"
	}

	return "Uknown"
}
