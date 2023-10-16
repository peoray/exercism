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
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}


// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    card1Value := ParseCard(card1)
    card2Value := ParseCard(card2)
    dealerCardValue := ParseCard(dealerCard)

	playerHandValue := card1Value + card2Value

    switch playerHandValue {
		case 22:
			return "P" 
        case 21:
    		if(dealerCardValue != 11 && dealerCardValue != 10) {
				return "W" 
            } else {
				return "S"
            }

        case 12, 13, 14, 15, 16:
    		if(dealerCardValue >= 7) {
				return "H"
			} else {
				return "S"
            }
    	case 17, 18, 19, 20:
			return "S"
		default:
			return "H"

    }
}
