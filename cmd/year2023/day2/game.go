package day2

import (
	"bytes"
	"fmt"
	"slices"
)

const (
	RedCubes   = 12
	GreenCubes = 13
	BlueCubes  = 14
)

type Game struct {
	ID     int
	Rounds []RGB
}

func (g *Game) UnmarshalText(text []byte) error {
	idSpec, roundsSpec, ok := bytes.Cut(text, []byte(":"))
	if !ok {
		return fmt.Errorf("invalid game: %s", text)
	}

	if _, err := fmt.Fscanf(bytes.NewReader(idSpec), "Game %d", &g.ID); err != nil {
		return err
	}

	roundsSpec = bytes.TrimSpace(roundsSpec)
	rounds := bytes.Split(roundsSpec, []byte(";"))
	g.Rounds = slices.Grow(g.Rounds, len(rounds))
	for i, roundSpec := range rounds {
		var round RGB
		if err := round.UnmarshalText(roundSpec); err != nil {
			return fmt.Errorf("failed to parse round %d: %w", i, err)
		}
		g.Rounds = append(g.Rounds, round)
	}

	return nil
}

func (g *Game) Valid() bool {
	for _, round := range g.Rounds {
		if !round.Valid() {
			return false
		}
	}
	return true
}

func (g *Game) Max() RGB {
	var result RGB
	for _, round := range g.Rounds {
		if result.R < round.R {
			result.R = round.R
		}
		if result.G < round.G {
			result.G = round.G
		}
		if result.B < round.B {
			result.B = round.B
		}
	}
	return result
}

func (g *Game) Power() int {
	m := g.Max()
	return m.R * m.G * m.B
}
