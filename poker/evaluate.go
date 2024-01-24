package poker

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"container/heap"

	"github.com/YoungsoonLee/poker/types"
)

// handCardCount represents the number of cards in a hand.
// It is a constant value for default.
const handCardCount = 5

// Hand represents the unique identifier of a hand.
// Cards represents the cards in the hand.
type Hand struct {
	HandID int
	Cards  []types.Card
}

// NewHand creates a new Hand with the specified handID.
// It initializes the Cards field with an empty slice of Card.
func NewHand(handID int) Hand {
	return Hand{
		HandID: handID,
		Cards:  make([]types.Card, handCardCount),
	}
}

// RandomCards generates a random Hand of cards based on the given handID.
// It uses a list of ranks and suits to randomly select a rank and suit for each card in the Hand.
// The generated Hand is returned.
func RandomCards(handID int) Hand {
	var ranks, suits []string
	for k := range types.RankMap {
		if k != "A" {
			ranks = append(ranks, k)
		}
	}

	for k := range types.SuitMap {
		suits = append(suits, k)
	}

	hand := NewHand(handID)

	for i := 0; i < handCardCount; i++ {
		// Create a new random number generator with a custom seed (e.g., current time)
		source := rand.NewSource(time.Now().UnixNano())
		rng := rand.New(source)

		rank := ranks[rng.Intn(len(ranks))]
		suit := suits[rng.Intn(len(suits))]

		hand.Cards[i] = types.Card{Rank: rank, Suit: suit}
	}

	return hand
}

// HasValidRanks checks if all the ranks in the hand are valid.
// It iterates through each card in the hand and checks if its rank exists in the RankMap.
// If any card has an invalid rank, it returns false. Otherwise, it returns true.
func (h Hand) HasValidRanks() bool {
	for _, card := range h.Cards {
		if _, ok := types.RankMap[card.Rank]; !ok {
			return false
		}
	}

	return true
}

// HasValidSuits checks if all the suits in the hand are valid.
// It iterates through each card in the hand and checks if its suit exists in the SuitMap.
// If any card has an invalid suit, it returns false.
// Otherwise, it returns true.
func (h Hand) HasValidSuits() bool {
	for _, card := range h.Cards {
		if _, ok := types.SuitMap[card.Suit]; !ok {
			return false
		}
	}

	return true
}

// ExtractRanksToInt extracts the ranks of the cards in the hand and returns them as a sorted slice of integers.
// The ranks are mapped to their corresponding integer values using the RankMap defined in the types package.
func (h Hand) ExtractRanksToInt() []int {
	var ranks []int

	for _, card := range h.Cards {
		ranks = append(ranks, types.RankMap[card.Rank])
	}

	sort.Ints(ranks)

	return ranks
}

// ExtractSuits extracts the suits from the hand and returns them as a sorted slice of strings.
// It iterates over each card in the hand and appends its suit to a string builder.
// The string builder is then split into individual characters and sorted alphabetically.
// The sorted suits are returned as a slice of strings.
func (h Hand) ExtractSuits() []string {

	var sb strings.Builder

	for _, card := range h.Cards {
		sb.WriteString(card.Suit)
	}

	r := strings.Split(sb.String(), "")

	sort.Strings(r)

	return r
}

// IsFlush checks if the hand is a flush, which means all cards have the same suit.
// It returns true if all cards have the same suit, and false otherwise.
func (h Hand) IsFlush() bool {
	suits := h.ExtractSuits()

	return suits[0] == suits[1] && suits[1] == suits[2] && suits[2] == suits[3] && suits[3] == suits[4]
}

// IsStraight checks if the hand represents a straight in poker.
// A straight is a hand where the ranks of the cards form a consecutive sequence.
// It also considers the special case of a low straight (A, 2, 3, 4, 5).
// Returns true if the hand is a straight, false otherwise.
func (h Hand) IsStraight() bool {
	ranks := h.ExtractRanksToInt()

	// check low straight
	if ranks[0] == 2 && ranks[1] == 3 && ranks[2] == 4 && ranks[3] == 5 && ranks[4] == 14 {
		return true
	}

	return ranks[0]+1 == ranks[1] && ranks[1]+1 == ranks[2] && ranks[2]+1 == ranks[3] && ranks[3]+1 == ranks[4]
}

// IsRoyalFlush checks if the hand is a royal flush.
// A royal flush is a hand that consists of the following cards in the same suit: 10, J, Q, K, A.
// It also checks if the hand is a flush.
// Returns true if the hand is a royal flush, false otherwise.
func (h Hand) IsRoyalFlush() bool {
	ranks := h.ExtractRanksToInt()

	return (ranks[0] == 10 && ranks[1] == 11 && ranks[2] == 12 && ranks[3] == 13 && ranks[4] == 14) && h.IsFlush()
}

// IsStraightFlush checks if the hand is a straight flush.
// A straight flush is a hand that is both a straight and a flush.
// Returns true if the hand is a straight flush, false otherwise.
func (h Hand) IsStraightFlush() bool {
	return h.IsStraight() && h.IsFlush()
}

// IsFourOfAKind checks if the hand contains four cards of the same rank.
// It returns true if the hand has four of a kind, otherwise it returns false.
func (h Hand) IsFourOfAKind() bool {
	ranks := h.ExtractRanksToInt()

	return (ranks[0] == ranks[1] && ranks[1] == ranks[2] && ranks[2] == ranks[3]) || (ranks[1] == ranks[2] && ranks[2] == ranks[3] && ranks[3] == ranks[4])
}

// IsThreeOfAKind checks if the hand contains three cards of the same rank.
// It returns true if the hand has three cards of the same rank, otherwise false.
func (h Hand) IsThreeOfAKind() bool {
	ranks := h.ExtractRanksToInt()

	return (ranks[0] == ranks[1] && ranks[1] == ranks[2]) || (ranks[1] == ranks[2] && ranks[2] == ranks[3]) || (ranks[2] == ranks[3] && ranks[3] == ranks[4])
}

// IsTwoPair checks if the hand contains two pairs of cards with the same rank.
// It returns true if the hand has two pairs, otherwise it returns false.
func (h Hand) IsTwoPair() bool {
	ranks := h.ExtractRanksToInt()

	return (ranks[0] == ranks[1] && ranks[2] == ranks[3]) || (ranks[0] == ranks[1] && ranks[3] == ranks[4]) || (ranks[1] == ranks[2] && ranks[3] == ranks[4])
}

// IsOnePair checks if the hand contains a pair of cards with the same rank.
// It returns true if a pair is found, otherwise it returns false.
func (h Hand) IsOnePair() bool {
	ranks := h.ExtractRanksToInt()

	return (ranks[0] == ranks[1]) || (ranks[1] == ranks[2]) || (ranks[2] == ranks[3]) || (ranks[3] == ranks[4])
}

// IsFullHouse checks if the hand is a full house.
// A full house is a hand that consists of three cards of the same rank and two cards of another rank.
// Returns true if the hand is a full house, false otherwise.
func (h Hand) IsFullHouse() bool {
	return h.IsThreeOfAKind() && h.IsTwoPair()
}

// HighCard returns the highest card rank in the hand.
// It extracts the ranks of the cards in the hand and returns the highest rank.
func (h Hand) HighCard() string {
	ranks := h.ExtractRanksToInt()

	return types.RankMapReverse[ranks[4]]
}

// Evaluate returns the name of the hand and its rank
func (h Hand) Evaluate() (string, int) {
	if h.IsRoyalFlush() {
		return "Royal Flush", 1
	} else if h.IsStraightFlush() {
		return "Straight Flush", 2
	} else if h.IsFourOfAKind() {
		return "Four of a Kind", 3
	} else if h.IsFullHouse() {
		return "Full House", 4
	} else if h.IsFlush() {
		return "Flush", 5
	} else if h.IsStraight() {
		return "Straight", 6
	} else if h.IsThreeOfAKind() {
		return "Three of a Kind", 7
	} else if h.IsTwoPair() {
		return "Two Pair", 8
	} else if h.IsOnePair() {
		return "One Pair", 9
	} else {
		return fmt.Sprintf("High Card - {%s}", h.HighCard()), 10
	}
}

// Hands represents a collection of Hand objects.
type Hands []Hand

// RandomCardsToHands generates a specified number of random hands.
// It takes an integer parameter 'num' representing the number of hands to generate.
// It returns a Hands object containing the generated hands.
func RandomCardsToHands(num int) Hands {
	if num < 1 {
		return nil
	}

	var hands Hands

	for i := 0; i < num; i++ {
		hands = append(hands, RandomCards(i))
	}

	return hands
}

// MinHeap is a type representing a minimum heap of Hand objects.
type MinHeap []Hand

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	_, r1 := h[i].Evaluate()
	_, r2 := h[j].Evaluate()

	return r1 < r2
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds a Hand to the MinHeap.
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Hand))
}

// Pop removes and returns the top element from the MinHeap.
// It modifies the underlying heap and returns the removed element.
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// HandResult is the unique identifier for a hand.
// Card is the list of cards in the hand.
// Rank is the rank of the hand.
// RankOrder is the order of the hand's rank.
type HandResult struct {
	HandID    int
	Card      []types.Card
	Rank      string
	RankOrder int
}

// EvaluateHands evaluates a collection of hands and returns the hand with the highest rank (the smallest number of rank).
// It's using minHeap to get the highest rank. the highest rank is the smallest number of rank.
func EvaluateHands(hands Hands) []HandResult {
	var minHeap MinHeap

	for _, hand := range hands {
		heap.Push(&minHeap, hand)
	}

	var handResults []HandResult

	for len(minHeap) > 0 {
		hand := heap.Pop(&minHeap).(Hand)

		rank, rankOrder := hand.Evaluate()
		handResults = append(handResults, HandResult{
			HandID:    hand.HandID,
			Card:      hand.Cards,
			Rank:      rank,
			RankOrder: rankOrder,
		})
	}

	return handResults
}
