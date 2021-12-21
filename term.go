package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type terminal struct {
	width  int
	height int
}

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
		case "down":
			snek.dir = DOWN
			snek.symbol = SNAKE_DOWN
		case "left":
			snek.dir = LEFT
			snek.symbol = SNAKE_LEFT
		case "right":
			snek.dir = RIGHT
			snek.symbol = SNAKE_RIGHT
		}

	case time.Time:
		snek.x += snek.dir.right
		snek.y += snek.dir.down
		return s, tick()
	}

	return s, nil

}

func (s snake) View() string {

	display := ""

	for i := 0; i < pty.height; i++ {

		if i == 0 {

			// Print Horizontal TOP Border
			for j := 0; j < pty.width; j++ {
				display += BORDER
			}
			display += "\n"

		} else if i == pty.height-1 {

			// Print Horizontal BOTTOM Border
			for j := 0; j < pty.width; j++ {
				display += BORDER
			}

		} else {

			// Print Vertical LEFT Border
			display += BORDER

			// Print Play Area
			for j := 0; j < pty.width-2; j++ {

				if i == snek.y && j == snek.x {
					display += snek.symbol
				} else if i == dinner.y && j == dinner.x {
					display += FOOD
				} else {
					display += " "
				}
			}

			// Print Vertical RIGHT Border
			display += BORDER
			display += "\n"

		}

	}

	return display
}
