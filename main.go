package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var deck [78]string

	// Major arcana
	deck[0] = "[0] The Fool"
	deck[1] = "[1] The Magician"
	deck[2] = "[2] The High Priestess"
	deck[3] = "[3] The Empress"
	deck[4] = "[4] The Emperor"
	deck[5] = "[5] The Hierophant"
	deck[6] = "[6] The Lovers"
	deck[7] = "[7] The Chariot"
	deck[8] = "[8] Strength"
	deck[9] = "[9] The Hermit"
	deck[10] = "[10] Wheel of Fortune"
	deck[11] = "[11] Justice"
	deck[12] = "[12] The Hanged Man"
	deck[13] = "[13] The Death"
	deck[14] = "[14] Temperance"
	deck[15] = "[15] Devil"
	deck[16] = "[16] The Tower"
	deck[17] = "[17] The Star"
	deck[18] = "[18] The Moon"
	deck[19] = "[19] The Sun"
	deck[20] = "[20] Judgement"
	deck[21] = "[21] The World"

	// Wands
	deck[22] = "Page of Wands"
	deck[23] = "Knight of Wands"
	deck[24] = "King of Wands"
	deck[25] = "Queen of Wands"
	deck[26] = "Ace of Wands"
	deck[27] = "Two of Wands"
	deck[28] = "Three of Wands"
	deck[29] = "Four of Wands"
	deck[30] = "Five of Wands"
	deck[31] = "Six of Wands"
	deck[32] = "Seven of Wands"
	deck[33] = "Eight of Wands"
	deck[34] = "Nine of Wands"
	deck[35] = "Ten of Wands"

	// Cups
	deck[36] = "Page of Cups"
	deck[37] = "Knight of Cups"
	deck[38] = "King of Cups"
	deck[39] = "Queen of Cups"
	deck[40] = "Ace of Cups"
	deck[41] = "Two of Cups"
	deck[42] = "Three of Cups"
	deck[43] = "Four of Cups"
	deck[44] = "Five of Cups"
	deck[45] = "Six of Cups"
	deck[46] = "Seven of Cups"
	deck[47] = "Eight of Cups"
	deck[48] = "Nine of Cups"
	deck[49] = "Ten of Cups"

	// Pentacles
	deck[50] = "Page of Pentacles"
	deck[51] = "Knight of Pentacles"
	deck[52] = "King of Pentacles"
	deck[53] = "Queen of Pentacles"
	deck[54] = "Ace of Pentacles"
	deck[55] = "Two of Pentacles"
	deck[56] = "Three of Pentacles"
	deck[57] = "Four of Pentacles"
	deck[58] = "Five of Pentacles"
	deck[59] = "Six of Pentacles"
	deck[60] = "Seven of Pentacles"
	deck[61] = "Eight of Pentacles"
	deck[62] = "Nine of Pentacles"
	deck[63] = "Ten of Pentacles"

	// Swords
	deck[64] = "Page of Swords"
	deck[65] = "Knight of Swords"
	deck[66] = "King of Swords"
	deck[67] = "Queen of Swords"
	deck[68] = "Ace of Swords"
	deck[69] = "Two of Swords"
	deck[70] = "Three of Swords"
	deck[71] = "Four of Swords"
	deck[72] = "Five of Swords"
	deck[73] = "Six of Swords"
	deck[74] = "Seven of Swords"
	deck[75] = "Eight of Swords"
	deck[76] = "Nine of Swords"
	deck[77] = "Ten of Swords"

	selectedCardIndex := rand.Intn(len(deck))
	selectedCard := deck[selectedCardIndex]

	fmt.Println(selectedCard)
}
