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

var area [][]uint8 // 0 -> Empty, 1 -> Snake Head, 2 -> Food, 3 -> Snake Body. Higher priority gets shown, except 0, which has least priority.
var snek snake
var dinner food

var HEIGHT int
var WIDTH int

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

	// Initialize play area
	HEIGHT = height - 3 // 2 for borders + 1 for Score
	WIDTH = width - 2 // 2 for borders
	area = make([][]uint8, HEIGHT) 
	for i:=0; i<HEIGHT; i++ {
		area[i] = make([]uint8, WIDTH)
	}

	// Initializing Snek
	snek.locationUpdate(WIDTH/2, HEIGHT/2)
	snek.lengthUpdate(1)
	snek.durationUpdate(400 * milli)
	snek.symbol = SNAKE_UP
	snek.dir = UP

	dinner.init() //Initialize food

	// Start interactive TUI
	if err := tea.NewProgram(snek).Start(); err != nil {
		fmt.Println("Oops, looks like there was an error", err)
		os.Exit(1)
	}
}
