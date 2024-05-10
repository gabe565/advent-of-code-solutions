// Code generated by "enumer -type Hand"; DO NOT EDIT.

package day7

import (
	"fmt"
	"strings"
)

const _HandName = "HighCardOnePairTwoPairThreeOfAKindFullHouseFourOfAKindFiveOfAKind"

var _HandIndex = [...]uint8{0, 8, 15, 22, 34, 43, 54, 65}

const _HandLowerName = "highcardonepairtwopairthreeofakindfullhousefourofakindfiveofakind"

func (i Hand) String() string {
	if i >= Hand(len(_HandIndex)-1) {
		return fmt.Sprintf("Hand(%d)", i)
	}
	return _HandName[_HandIndex[i]:_HandIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _HandNoOp() {
	var x [1]struct{}
	_ = x[HighCard-(0)]
	_ = x[OnePair-(1)]
	_ = x[TwoPair-(2)]
	_ = x[ThreeOfAKind-(3)]
	_ = x[FullHouse-(4)]
	_ = x[FourOfAKind-(5)]
	_ = x[FiveOfAKind-(6)]
}

var _HandValues = []Hand{HighCard, OnePair, TwoPair, ThreeOfAKind, FullHouse, FourOfAKind, FiveOfAKind}

var _HandNameToValueMap = map[string]Hand{
	_HandName[0:8]:        HighCard,
	_HandLowerName[0:8]:   HighCard,
	_HandName[8:15]:       OnePair,
	_HandLowerName[8:15]:  OnePair,
	_HandName[15:22]:      TwoPair,
	_HandLowerName[15:22]: TwoPair,
	_HandName[22:34]:      ThreeOfAKind,
	_HandLowerName[22:34]: ThreeOfAKind,
	_HandName[34:43]:      FullHouse,
	_HandLowerName[34:43]: FullHouse,
	_HandName[43:54]:      FourOfAKind,
	_HandLowerName[43:54]: FourOfAKind,
	_HandName[54:65]:      FiveOfAKind,
	_HandLowerName[54:65]: FiveOfAKind,
}

var _HandNames = []string{
	_HandName[0:8],
	_HandName[8:15],
	_HandName[15:22],
	_HandName[22:34],
	_HandName[34:43],
	_HandName[43:54],
	_HandName[54:65],
}

// HandString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func HandString(s string) (Hand, error) {
	if val, ok := _HandNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _HandNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Hand values", s)
}

// HandValues returns all values of the enum
func HandValues() []Hand {
	return _HandValues
}

// HandStrings returns a slice of all String values of the enum
func HandStrings() []string {
	strs := make([]string, len(_HandNames))
	copy(strs, _HandNames)
	return strs
}

// IsAHand returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Hand) IsAHand() bool {
	for _, v := range _HandValues {
		if i == v {
			return true
		}
	}
	return false
}