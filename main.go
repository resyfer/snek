package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

var milli = time.Millisecond
var random *rand.Rand

var snek snake
var pty terminal
var dinner food

func main() {

	// Setting random's seed
	seed := rand.NewSource(time.Now().UnixNano())
	random = rand.New(seed)

	// Getting terminal height
	width, height, err := term.GetSize(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pty.width = width
	pty.height = height - 1 // For score

	// Initializing Snek
	snek.locationUpdate(width/2, height/2)
	snek.lengthUpdate(1)
	snek.durationUpdate(500 * milli)
	snek.symbol = SNAKE_RIGHT
	snek.dir = RIGHT

	dinner.init()

	if err := tea.NewProgram(snek).Start(); err != nil {
		fmt.Println("Oops, looks like there was an error", err)
		os.Exit(1)
	}
}
