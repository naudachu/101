package main

import (
	"101/player"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

func main() {
	players := gameInit()
	gameLoop(players)
}

func gameLoop(players []*player.Player) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {

			str := strings.Split(text, " ")
			switch str[0] {
			case "add":
				if len(str) != 3 {
					log.Println("wrong command")
					break
				}
				i, err := strconv.Atoi(str[1])
				if err != nil {
					log.Println("player id parse error")
				}

				players[i].SetPoints(str[2])
				printPlayers(players)

			case "sub":
				i, err := strconv.Atoi(str[1])
				if err != nil {
					log.Println("player id parse error")
				}

				// Check if -points can be converted to int
				subPoints, err := strconv.Atoi(str[2])
				if err != nil {
					log.Println("subPoints convertion to int failed")
				}

				players[i].SubPoints(subPoints)
				printPlayers(players)

			default:
				log.Println("wrong command")
			}

		} else {
			// exit if user entered an empty string
			break
		}

	}

	// handle error

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}

func gameInit() []*player.Player {
	clear()
	flag.Parse()
	names := flag.Args()

	var players []*player.Player

	switch len(names) {
	case 0:
		players = append(players, player.NewPlayer("Player 1"))
		players = append(players, player.NewPlayer("Player 2"))
	default:
		for _, e := range names {
			{
				players = append(players, player.NewPlayer(e))
			}
		}
	}

	printPlayers(players)
	return players
}

func printPlayers(list []*player.Player) {
	clear()
	for i, p := range list {
		color.HEX(p.Color()).Printf("%d. %s\n", i, p.Name())
		fmt.Println(p.Score())
	}
}

func clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
