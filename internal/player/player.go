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
		str = append(str, fmt.Sprint(e))
	}
	if str == nil {
		str = append(str, "-")
	}
	scoreString := strings.Join(str, " ")

	return fmt.Sprintf("%d | %s", p.score, scoreString)
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

func (p *Player) SubPoints(collected int) *Player {
	if collected < 0 {
		p.rounds = append(p.rounds, collected)
		p.score = p.score + collected
	} else {
		p.rounds = append(p.rounds, -collected)
		p.score = p.score - collected
	}
	if p.score < 0 {
		p.score = 0
	}
	return p
}

func (p *Player) Win() *Player {
	p.rounds = []int{}
	p.score = 0
	return p
}