package poker

import (
	"reflect"
	"testing"

	"github.com/YoungsoonLee/poker/types"
)

func TestHand_HasValidSuits(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want bool
	}{
		{
			name: "right suit",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "S"},
					{Rank: "4", Suit: "S"},
					{Rank: "5", Suit: "S"},
					{Rank: "6", Suit: "S"},
				},
			},

			want: true,
		},
		{
			name: "wrong suit",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "H"},
					{Rank: "4", Suit: "D"},
					{Rank: "5", Suit: "C"},
					{Rank: "6", Suit: "B"},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.HasValidSuits(); got != tt.want {
				t.Errorf("Hand.HasValidSuits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_HasValidRanks(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want bool
	}{
		{
			name: "right rank number",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "S"},
					{Rank: "4", Suit: "S"},
					{Rank: "5", Suit: "S"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: true,
		},
		{
			name: "right rank letter",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "K", Suit: "S"},
					{Rank: "J", Suit: "S"},
					{Rank: "Q", Suit: "S"},
					{Rank: "A", Suit: "S"},
				},
			},
			want: true,
		},
		{
			name: "wrong rank number",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "S"},
					{Rank: "4", Suit: "S"},
					{Rank: "5", Suit: "S"},
					{Rank: "14", Suit: "S"},
				},
			},
			want: false,
		},
		{
			name: "wrong rank letter",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "K", Suit: "S"},
					{Rank: "J", Suit: "S"},
					{Rank: "Q", Suit: "S"},
					{Rank: "B", Suit: "S"},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.HasValidRanks(); got != tt.want {
				t.Errorf("Hand.HasValidRanks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_ExtractRanksToInt(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want []int
	}{
		{
			name: "number",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "S"},
					{Rank: "4", Suit: "S"},
					{Rank: "5", Suit: "S"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: []int{2, 3, 4, 5, 6},
		},
		{
			name: "has same rank",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "2", Suit: "S"},
					{Rank: "4", Suit: "S"},
					{Rank: "5", Suit: "S"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: []int{2, 2, 4, 5, 6},
		},
		{
			name: "number and letter",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "K", Suit: "S"},
					{Rank: "J", Suit: "S"},
					{Rank: "Q", Suit: "S"},
					{Rank: "A", Suit: "S"},
				},
			},
			want: []int{2, 11, 12, 13, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.ExtractRanksToInt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hand.ExtractRanksToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_ExtractSuits(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want []string
	}{
		{
			name: "number",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "5", Suit: "H"},
					{Rank: "5", Suit: "D"},
				},
			},
			want: []string{"C", "D", "D", "H", "S"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.ExtractSuits(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hand.ExtractSuits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsFlush(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want bool
	}{
		{
			name: "flush",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "S"},
					{Rank: "4", Suit: "S"},
					{Rank: "5", Suit: "S"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: true,
		},
		{
			name: "not flush",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "S"},
					{Rank: "4", Suit: "S"},
					{Rank: "5", Suit: "S"},
					{Rank: "6", Suit: "H"},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.IsFlush(); got != tt.want {
				t.Errorf("Hand.IsFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsStraight(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want bool
	}{
		{
			name: "straight",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "5", Suit: "H"},
					{Rank: "6", Suit: "D"},
				},
			},
			want: true,
		},
		{
			name: "not straight",
			h: Hand{
				Cards: []types.Card{
					{Rank: "3", Suit: "S"},
					{Rank: "2", Suit: "S"},
					{Rank: "4", Suit: "S"},
					{Rank: "5", Suit: "S"},
					{Rank: "7", Suit: "H"},
				},
			},
			want: false,
		},
		{
			name: "low straight",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "S"},
					{Rank: "A", Suit: "S"},
					{Rank: "5", Suit: "S"},
					{Rank: "4", Suit: "H"},
				},
			},
			want: true,
		},
		{
			name: "high straight",
			h: Hand{
				Cards: []types.Card{
					{Rank: "T", Suit: "S"},
					{Rank: "A", Suit: "S"},
					{Rank: "Q", Suit: "S"},
					{Rank: "K", Suit: "S"},
					{Rank: "J", Suit: "H"},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.IsStraight(); got != tt.want {
				t.Errorf("Hand.IsStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsRoyalFlush(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want bool
	}{
		{
			name: "royal flush",
			h: Hand{
				Cards: []types.Card{
					{Rank: "T", Suit: "S"},
					{Rank: "A", Suit: "S"},
					{Rank: "Q", Suit: "S"},
					{Rank: "K", Suit: "S"},
					{Rank: "J", Suit: "S"},
				},
			},
			want: true,
		},
		{
			name: "not royal flush",
			h: Hand{
				Cards: []types.Card{
					{Rank: "T", Suit: "S"},
					{Rank: "A", Suit: "S"},
					{Rank: "Q", Suit: "S"},
					{Rank: "K", Suit: "S"},
					{Rank: "2", Suit: "S"},
				},
			},
			want: false,
		},
		{
			name: "not royal flush",
			h: Hand{
				Cards: []types.Card{
					{Rank: "T", Suit: "S"},
					{Rank: "A", Suit: "S"},
					{Rank: "Q", Suit: "S"},
					{Rank: "K", Suit: "S"},
					{Rank: "J", Suit: "H"},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.IsRoyalFlush(); got != tt.want {
				t.Errorf("Hand.IsRoyalFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsStraightFlush(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want bool
	}{
		{
			name: "straight flush",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "S"},
					{Rank: "4", Suit: "S"},
					{Rank: "5", Suit: "S"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: true,
		},
		{
			name: "not straight flush",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "S"},
					{Rank: "4", Suit: "S"},
					{Rank: "5", Suit: "S"},
					{Rank: "6", Suit: "H"},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.IsStraightFlush(); got != tt.want {
				t.Errorf("Hand.IsStraightFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsFourOfAKind(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want bool
	}{
		{
			name: "four of a kind",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "2", Suit: "D"},
					{Rank: "2", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: true,
		},
		{
			name: "not four of a kind",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "2", Suit: "D"},
					{Rank: "3", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: false,
		},
		{
			name: "four of a kind",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "J", Suit: "H"},
					{Rank: "J", Suit: "D"},
					{Rank: "J", Suit: "C"},
					{Rank: "J", Suit: "S"},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.IsFourOfAKind(); got != tt.want {
				t.Errorf("Hand.IsFourOfAKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsThreeOfAKind(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want bool
	}{
		{
			name: "three of a kind",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "2", Suit: "D"},
					{Rank: "3", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: true,
		},
		{
			name: "not three of a kind",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "3", Suit: "D"},
					{Rank: "3", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.IsThreeOfAKind(); got != tt.want {
				t.Errorf("Hand.IsThreeOfAKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsTwoPair(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want bool
	}{
		{
			name: "two pair",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "3", Suit: "D"},
					{Rank: "3", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: true,
		},
		{
			name: "not two pair",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "3", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: false,
		},
		{
			name: "two pair",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "J", Suit: "H"},
					{Rank: "J", Suit: "D"},
					{Rank: "Q", Suit: "C"},
					{Rank: "Q", Suit: "S"},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.IsTwoPair(); got != tt.want {
				t.Errorf("Hand.IsTwoPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsOnePair(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want bool
	}{
		{
			name: "one pair",
			h: Hand{
				Cards: []types.Card{
					{Rank: "5", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "3", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "5", Suit: "S"},
				},
			},
			want: true,
		},
		{
			name: "not one pair",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "H"},
					{Rank: "4", Suit: "D"},
					{Rank: "5", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.IsOnePair(); got != tt.want {
				t.Errorf("Hand.IsOnePair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsFullHouse(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want bool
	}{
		{
			name: "full house",
			h: Hand{
				Cards: []types.Card{
					{Rank: "5", Suit: "S"},
					{Rank: "5", Suit: "H"},
					{Rank: "5", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "4", Suit: "S"},
				},
			},
			want: true,
		},
		{
			name: "not full house",
			h: Hand{
				Cards: []types.Card{
					{Rank: "5", Suit: "S"},
					{Rank: "5", Suit: "H"},
					{Rank: "5", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.IsFullHouse(); got != tt.want {
				t.Errorf("Hand.IsFullHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsHighCard(t *testing.T) {
	tests := []struct {
		name string
		h    Hand
		want string
	}{
		{
			name: "high card",
			h: Hand{
				Cards: []types.Card{
					{Rank: "5", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "3", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want: "6",
		},
		{
			name: "high card",
			h: Hand{
				Cards: []types.Card{
					{Rank: "5", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "3", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "T", Suit: "S"},
				},
			},
			want: "T",
		},
		{
			name: "high card",
			h: Hand{
				Cards: []types.Card{
					{Rank: "5", Suit: "S"},
					{Rank: "Q", Suit: "H"},
					{Rank: "3", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "T", Suit: "S"},
				},
			},
			want: "Q",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.HighCard(); got != tt.want {
				t.Errorf("Hand.HighCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_Evaluate(t *testing.T) {
	tests := []struct {
		name     string
		h        Hand
		want     string
		wantRank int
	}{
		{
			name: "royal flush",
			h: Hand{
				Cards: []types.Card{
					{Rank: "T", Suit: "S"},
					{Rank: "A", Suit: "S"},
					{Rank: "Q", Suit: "S"},
					{Rank: "K", Suit: "S"},
					{Rank: "J", Suit: "S"},
				},
			},
			want:     "Royal Flush",
			wantRank: 1,
		},
		{
			name: "straight flush",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "S"},
					{Rank: "4", Suit: "S"},
					{Rank: "5", Suit: "S"},
					{Rank: "6", Suit: "S"},
				},
			},
			want:     "Straight Flush",
			wantRank: 2,
		},
		{
			name: "four of a kind",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "2", Suit: "D"},
					{Rank: "2", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want:     "Four of a Kind",
			wantRank: 3,
		},
		{
			name: "full house",
			h: Hand{
				Cards: []types.Card{
					{Rank: "5", Suit: "S"},
					{Rank: "5", Suit: "H"},
					{Rank: "5", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "4", Suit: "S"},
				},
			},
			want:     "Full House",
			wantRank: 4,
		},
		{
			name: "flush",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "S"},
					{Rank: "4", Suit: "S"},
					{Rank: "8", Suit: "S"},
					{Rank: "6", Suit: "S"},
				},
			},
			want:     "Flush",
			wantRank: 5,
		},
		{
			name: "straight",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "3", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "5", Suit: "H"},
					{Rank: "6", Suit: "D"},
				},
			},
			want:     "Straight",
			wantRank: 6,
		},

		{
			name: "three of a kind",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "2", Suit: "D"},
					{Rank: "3", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want:     "Three of a Kind",
			wantRank: 7,
		},
		{
			name: "two pair",
			h: Hand{
				Cards: []types.Card{
					{Rank: "2", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "3", Suit: "D"},
					{Rank: "3", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want:     "Two Pair",
			wantRank: 8,
		},
		{
			name: "one pair",
			h: Hand{
				Cards: []types.Card{
					{Rank: "5", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "3", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "5", Suit: "S"},
				},
			},
			want:     "One Pair",
			wantRank: 9,
		},

		{
			name: "high card",
			h: Hand{
				Cards: []types.Card{
					{Rank: "5", Suit: "S"},
					{Rank: "2", Suit: "H"},
					{Rank: "9", Suit: "D"},
					{Rank: "4", Suit: "C"},
					{Rank: "6", Suit: "S"},
				},
			},
			want:     "High Card - {9}",
			wantRank: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, gotInt := tt.h.Evaluate(); got != tt.want || gotInt != tt.wantRank {
				t.Errorf("Hand.Rank() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestHand_RandomCardsToHands(t *testing.T) {

	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "test case 1",
			num:  1,
			want: 1,
		},
		{
			name: "test case 6",
			num:  6,
			want: 6,
		},
		{
			name: "test case 0",
			num:  0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomCardsToHands(tt.num)
			if len(got) != tt.want {
				t.Errorf("RandomCardsToHands() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

// TODO: need to fix
func TestHand_EvaluateHands(t *testing.T) {
	hands := []Hand{
		{
			HandID: 1,
			// straight flush, 2.
			Cards: []types.Card{
				{Rank: "2", Suit: "S"},
				{Rank: "3", Suit: "S"},
				{Rank: "4", Suit: "S"},
				{Rank: "5", Suit: "S"},
				{Rank: "6", Suit: "S"},
			},
		},
		{
			HandID: 2,
			// royal straight, 6.
			Cards: []types.Card{
				{Rank: "A", Suit: "S"},
				{Rank: "K", Suit: "D"},
				{Rank: "Q", Suit: "H"},
				{Rank: "J", Suit: "C"},
				{Rank: "T", Suit: "S"},
			},
		},
		{
			// flush, 5.
			HandID: 3,
			Cards: []types.Card{
				{Rank: "2", Suit: "H"},
				{Rank: "8", Suit: "H"},
				{Rank: "4", Suit: "H"},
				{Rank: "5", Suit: "H"},
				{Rank: "6", Suit: "H"},
			},
		},
		{
			// one pair, 9.
			HandID: 4,
			Cards: []types.Card{
				{Rank: "2", Suit: "H"},
				{Rank: "8", Suit: "D"},
				{Rank: "8", Suit: "C"},
				{Rank: "5", Suit: "H"},
				{Rank: "6", Suit: "K"},
			},
		},
		{
			// two pair, 8.
			HandID: 5,
			Cards: []types.Card{
				{Rank: "2", Suit: "H"},
				{Rank: "8", Suit: "D"},
				{Rank: "8", Suit: "C"},
				{Rank: "5", Suit: "H"},
				{Rank: "5", Suit: "K"},
			},
		},
		{
			// High Card, 10.
			HandID: 6,
			Cards: []types.Card{
				{Rank: "2", Suit: "H"},
				{Rank: "8", Suit: "D"},
				{Rank: "J", Suit: "C"},
				{Rank: "5", Suit: "H"},
				{Rank: "A", Suit: "K"},
			},
		},
	}

	result := EvaluateHands(hands)

	for i, v := range result {
		if i == 0 {
			if v.HandID != 1 {
				t.Errorf("EvaluateHands() = %v, want %v", v.HandID, 1)
			}
		} else if i == 1 {
			if v.HandID != 3 {
				t.Errorf("EvaluateHands() = %v, want %v", v.HandID, 3)
			}
		} else if i == 2 {
			if v.HandID != 2 {
				t.Errorf("EvaluateHands() = %v, want %v", v.HandID, 2)
			}
		} else if i == 3 {
			if v.HandID != 5 {
				t.Errorf("EvaluateHands() = %v, want %v", v.HandID, 5)
			}
		} else if i == 4 {
			if v.HandID != 4 {
				t.Errorf("EvaluateHands() = %v, want %v", v.HandID, 4)
			}
		} else if i == 5 {
			if v.HandID != 6 {
				t.Errorf("EvaluateHands() = %v, want %v", v.HandID, 6)
			}
		}

	}
}
