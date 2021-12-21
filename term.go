package main

import (
	"fmt"
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
		snek.x += snek.dir.right
		snek.y += snek.dir.down

		if snek.foodCollision() {
			dinner.init()
			snek.points++
			
			if snek.dur > 100 * time.Millisecond {
				snek.dur -= 50 * time.Millisecond
			}
		}

		return s, tick()
	}

	return s, nil

}

func (s snake) View() string {

	display := ""

	display += fmt.Sprintf("Points Scored : %d\n", snek.points)

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
