package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func tick() tea.Cmd {

	return tea.Tick(snek.dur, func(t time.Time) tea.Msg {
		return t
	})

}

func (s snake) Init() tea.Cmd {
	return tick()
}

func (s snake) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return s, tea.Quit

		case "up":
			snek.dir = UP
			snek.symbol = SNAKE_UP
			return s, nil
		case "down":
			snek.dir = DOWN
			snek.symbol = SNAKE_DOWN
			return s, nil
		case "left":
			snek.dir = LEFT
			snek.symbol = SNAKE_LEFT
			return s, nil
		case "right":
			snek.dir = RIGHT
			snek.symbol = SNAKE_RIGHT
			return s, nil
		}

	case time.Time:


		return s, tick()
	}

	return s, nil

}

func (s snake) View() string {

	display := ""

	display += fmt.Sprintf("Points Scored : %v\n", snek.body)


	// Print Horizontal TOP Border
	display += BORDER
	for i := 0; i < WIDTH; i++ {
		display += BORDER
	}
	display += BORDER
	display += "\n"

	for i := 0; i < HEIGHT; i++ {

		// Print Vertical LEFT Border
		display += BORDER

		// Print Play Area
		for i := 0; i < WIDTH; i++ {
			display += " "
		}

		// Print Vertical RIGHT Border
		display += BORDER
		display += "\n"
	}

	// Print Horizontal BOTTOM Border
	display += BORDER
	for i := 0; i < WIDTH; i++ {
		display += BORDER
	}
	display += BORDER

	return display
}
