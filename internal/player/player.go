package player

import (
	"101/internal/clc"
	"fmt"
	"strings"
)

type Player struct {
	name    string
	rounds  []int // A number of points calculated at the each round;
	score   int   // Total points;
	color   string
	strikes int
}

func (p *Player) Title() string {
	var xx string
	for x := 0; x < p.strikes; x++ {
		xx += "•"
	}
	return p.name + " " + xx
}

func (p *Player) Description() string {
	var str []string
	for _, e := range p.rounds {
		switch len(fmt.Sprint(e)) {
		case 0:
			str = append(str, "-")
		case 1:
			str = append(str, "  ", fmt.Sprint(e))
		case 2:
			str = append(str, " ", fmt.Sprint(e))
		default:
			str = append(str, "", fmt.Sprint(e))
		}
	}

	var scoreString string

	switch len(fmt.Sprint(p.score)) {
	case 1:
		scoreString = fmt.Sprint("  ", p.score)
	case 2:
		scoreString = fmt.Sprint(" ", p.score)
	default:
		scoreString = fmt.Sprint(p.score)
	}

	roundsString := strings.Join(str, " ")

	return fmt.Sprintf("%s | %s", scoreString, roundsString)
}

func (p *Player) Color() string {
	return p.color
}

// NewPlayer method creates a New Player object;
func NewPlayer(name string) *Player {
	p := &Player{
		name:   name,
		rounds: []int{},
		score:  0,
	}
	return p
}

// SetPoints method parse cards string and calcs points amount;
// empty string should add ZERO points to the Player's score,
// valid card string should be calced with 'clc' package;
func (p *Player) SetPoints(cards string) *Player {
	switch cards {

	case "":
		// [ ] hm?

	default:
		collected := clc.Add(cards)
		p.rounds = append(p.rounds, collected)

		if p.score+collected >= 101 {
			p.strikes += 1
			p.score = 0
			p.rounds = []int{}
		} else {
			p.score = p.score + collected
		}
	}

	return p
}

func (p *Player) SubPoints(collected string) *Player {
	switch collected {
	case "q":
		p.rounds = append(p.rounds, -20)
		if p.score-20 < 0 {
			p.score = 0
		} else {
			p.score = p.score - 20
		}

	case "qq":
		p.rounds = append(p.rounds, -40)
		if p.score-40 < 0 {
			p.score = 0
		} else {
			p.score = p.score - 40
		}
	}
	return p
}

func (p *Player) Win() *Player {
	p.rounds = []int{}
	p.score = 0
	return p
}
