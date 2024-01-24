// Package types provides types and functions related to playing cards.
package types

import (
	"fmt"
	"strings"
)

// RankMap maps the rank of a card to its corresponding integer value.
var RankMap = map[string]int{
	"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
}

// RankMapReverse maps the integer value of a card to its corresponding rank.
var RankMapReverse = map[int]string{
	2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "T", 11: "J", 12: "Q", 13: "K", 14: "A",
}

// SuitMap maps the suit of a card to an empty struct.
var SuitMap = map[string]struct{}{
	"S": {}, "H": {}, "D": {}, "C": {},
}

// Card represents a playing card.
type Card struct {
	Rank string // Rank of the card (2, 3, 4, 5, 6, 7, 8, 9, T, J, Q, K, A)
	Suit string // Suit of the card (S, H, D, C)
}

// NewCard creates a new Card object based on the provided inputCard string.
// The inputCard string should be in the format "3s4h5d6c7s" or "9H3CTSQSAS",
// where odd indices represent the rank and even indices represent the suit.
// The function returns the created Card object and an error if the inputCard is invalid.
func NewCard(inputCard string) ([]Card, error) {
	if len(inputCard) != 10 {
		return []Card{}, fmt.Errorf("invalid card string: %s. card string should be like this. ex) 3s4h5d6c7s or 9H3CTSQSAS", inputCard)
	}

	rank := make([]string, 0)
	suit := make([]string, 0)
	card := make([]Card, 0)

	// even index is rank, odd index is suit
	for i, c := range inputCard {
		if i%2 == 0 {
			upper := strings.ToUpper(string(c))
			if _, ok := RankMap[upper]; !ok {
				return []Card{}, fmt.Errorf("invalid card string: %s. rank should be 2,3,4,5,6,7,8,9,T,J,Q,K,A", inputCard)
			}
			rank = append(rank, upper)
		} else {
			upper := strings.ToUpper(string(c))
			if _, ok := SuitMap[upper]; !ok {
				return []Card{}, fmt.Errorf("invalid card string: %s. suit should be S,H,D,C", inputCard)
			}
			suit = append(suit, upper)
		}
	}

	for i := 0; i < len(rank); i++ {
		card = append(card, Card{Rank: rank[i], Suit: suit[i]})
	}

	return card, nil
}

// String returns a string representation of the card.
func (c Card) String() string {
	return c.Rank + c.Suit
}
