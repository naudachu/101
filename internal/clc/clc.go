package clc

import "strings"

var CN = map[string]int{
	"j": 2,
	"q": 3,
	"k": 4,
	"t": 11,
	"0": 10,
	"9": 0,
}

func Add(in string) int {
	in = strings.ToLower(in)
	var output int
	switch in {
	case "qqqqq":
		output = 100
	case "qqqq":
		output = 80
	case "qqq":
		output = 60
	case "qq":
		output = 40
	case "q":
		output = 20
	default:
		for _, ch := range in {
			output += add(ch)
		}
	}

	return output
}

func add(ch rune) int {
	i, ok := CN[string(ch)]
	if ok {
		return i
	}

	return int(ch) - '0'
}
