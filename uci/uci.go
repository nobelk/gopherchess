package uci

import (
	"fmt"
	"github.com/nobelk/gopherchess/engine"
	"strings"
)

func ExecuteUCIcommand(engine engine.Engine, input string) {
	if input == "" || input == "null" {
		return
	}

	tokens := strings.Fields(strings.TrimSpace(input))

	switch tokens[0] {
	case "uci":
		fmt.Println("uciok")
	case "isready":
		fmt.Println("readyok")
	case "position":
		fmt.Println("position")
	case "go":
		fmt.Println("ucigo")
	case "ucinewgame":
		engine.Stop()
		engine.Reset()
	case "stop":
		engine.Stop()
	case "quit":
		engine.Quit()
	case "setoption":
		fmt.Println("UciSetOption")
	default:
		fmt.Println("Unknown command:" + tokens[0])
	}
}
