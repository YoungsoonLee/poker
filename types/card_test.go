package types

import (
	"reflect"
	"testing"
)

// TestNewCard tests the NewCard function.
func TestNewCard(t *testing.T) {
	type args struct {
		inputCard string
	}
	tests := []struct {
		name    string
		args    args
		want    []Card
		wantErr bool
	}{
		{
			name: "valid card string",
			args: args{
				inputCard: "3s4h5d6c7s",
			},
			want: []Card{
				{Rank: "3", Suit: "S"},
				{Rank: "4", Suit: "H"},
				{Rank: "5", Suit: "D"},
				{Rank: "6", Suit: "C"},
				{Rank: "7", Suit: "S"},
			},
			wantErr: false,
		},
		{
			name: "invalid card string",
			args: args{
				inputCard: "3s4h5d6c7",
			},
			want:    []Card{},
			wantErr: true,
		},
		{
			name: "invalid card string",
			args: args{
				inputCard: "2p4h5d6c7",
			},
			want:    []Card{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCard(tt.args.inputCard)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
