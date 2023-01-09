package player

import (
	"101/clc"
	"fmt"
	"strings"
)

var number int = 0
var colors = []string{
	"#a83291",
	"#f5428d",
	"#5142f5",
	"#42f59e",
	"#8df542",
	"#f5e642",
}

type Player struct {
	name   string
	rounds []int // A number of points calculated at the each round;
	score  int   // Total points;
	color  string
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Score() string {
	// Rounds to string
	var str []string
	for _, e := range p.rounds {
		str = append(str, fmt.Sprint(e))
	}
	if str == nil {
		str = append(str, "-")
	}
	scoreString := strings.Join(str, " ")

	return fmt.Sprintf("%d | %s\n", p.score, scoreString)
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
		color:  colors[number],
	}
	number += 1
	return p
}

// SetPoints method parse cards string and calcs points amount;
// empty string should add ZERO points to the Player's score,
// valid card string should be calced with 'clc' package;
func (p *Player) SetPoints(cards string) *Player {
	switch cards {

	case "":
		p.rounds = append(p.rounds, 0)

	default:
		collected := clc.Add(cards)
		p.rounds = append(p.rounds, collected)
		p.score = p.score + collected
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
