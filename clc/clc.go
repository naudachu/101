package clc

var CN = map[string]int{
	"j": 2,
	"q": 3,
	"k": 4,
	"t": 11,
	"0": 10,
	"9": 0,
}

func Add(in string) int {
	var output int
	for _, ch := range in {
		output += add(ch)
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
