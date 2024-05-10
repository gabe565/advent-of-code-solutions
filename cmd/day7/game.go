package day7

import (
	"bufio"
	"bytes"
	"cmp"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"

	"github.com/gabe565/advent-of-code-2023/internal/util"
)

type Game struct {
	Rounds []Round
}

func (g *Game) Decode(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var round Round
		if err := round.UnmarshalText(scanner.Bytes()); err != nil {
			return err
		}
		g.Rounds = append(g.Rounds, round)
	}
	return scanner.Err()
}

func (g *Game) Winnings(wildcard bool) int {
	rounds := slices.Clone(g.Rounds)

	slices.SortStableFunc(rounds, func(a, b Round) int {
		if v := cmp.Compare(a.Hand(wildcard), b.Hand(wildcard)); v != 0 {
			return v
		}
		for i := range a.Cards {
			if v := cmp.Compare(rank(a.Cards[i], wildcard), rank(b.Cards[i], wildcard)); v != 0 {
				return v
			}
		}
		return 0
	})

	var winnings int
	for i, round := range rounds {
		winnings += round.Bid * (i + 1)
	}
	return winnings
}

type Round struct {
	Cards []byte
	Bid   int
}

var ErrInvalidInput = errors.New("invalid input")

func (r *Round) UnmarshalText(text []byte) error {
	var bid []byte
	var ok bool
	r.Cards, bid, ok = bytes.Cut(slices.Clone(text), []byte(" "))
	if !ok {
		return fmt.Errorf("%w: %s", ErrInvalidInput, string(text))
	}

	var err error
	r.Bid, err = strconv.Atoi(string(bid))
	if err != nil {
		return err
	}

	return nil
}

func (r *Round) Hand(wildcard bool) Hand {
	var best Hand
	for _, hand := range HandValues() {
		if hand.Matches(r.Cards, wildcard) {
			best = hand
		}
	}
	return best
}

//go:generate enumer -type Hand

type Hand uint16

const (
	HighCard Hand = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (h Hand) Matches(cards []byte, wildcard bool) bool {
	cards = slices.Clone(cards)
	slices.Sort(cards)
	uniq := slices.Compact(slices.Clone(cards))
	if wildcard && len(uniq) != 1 {
		for i, b := range uniq {
			if b == 'J' {
				uniq = slices.Delete(uniq, i, i+1)
				break
			}
		}
	}

	switch h {
	case FiveOfAKind:
		return len(uniq) == 1
	case FourOfAKind:
		if len(uniq) == 2 {
			counts := getCounts(cards, uniq, wildcard)
			return counts[0] == 1 && counts[1] == 4
		}
	case FullHouse:
		if len(uniq) == 2 {
			counts := getCounts(cards, uniq, wildcard)
			return counts[0] == 2 && counts[1] == 3
		}
	case ThreeOfAKind:
		if len(uniq) == 3 {
			counts := getCounts(cards, uniq, wildcard)
			return counts[0] == 1 && counts[1] == 1 && counts[2] == 3
		}
	case TwoPair:
		if len(uniq) == 3 {
			counts := getCounts(cards, uniq, wildcard)
			return counts[0] == 1 && counts[1] == 2 && counts[2] == 2
		}
	case OnePair:
		return len(uniq) == 4
	case HighCard:
		return len(uniq) == 5
	}
	return false
}

func getCounts(cards, uniq []byte, wildcard bool) []int {
	counts := make([]int, 0, len(uniq))
	for _, v := range uniq {
		counts = append(counts, util.Count(cards, v))
	}
	slices.Sort(counts)
	if wildcard {
		counts[len(counts)-1] += util.Count(cards, 'J')
	}
	return counts
}

func rank(b byte, wildcard bool) int {
	switch b {
	case '2':
		return 1
	case '3':
		return 2
	case '4':
		return 3
	case '5':
		return 4
	case '6':
		return 5
	case '7':
		return 6
	case '8':
		return 7
	case '9':
		return 8
	case 'T':
		return 9
	case 'J':
		if !wildcard {
			return 10
		}
	case 'Q':
		return 11
	case 'K':
		return 12
	case 'A':
		return 13
	}
	return 0
}
