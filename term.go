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

		if snek.boundaryCollision() {
			return s, tea.Quit
		}

		if snek.foodCollision() {
			snek.pointsInr()
			snek.lenInr()
			snek.bodyPush(snek.x, snek.y)
			snek.move()
			dinner.init()

			if snek.dur >= 100*milli {
				snek.durationUpdate(snek.dur - 50*milli)
			}
		}

		snek.bodyPush(snek.x, snek.y)
		snek.move()
		snek.bodyPop()

		return s, tick()
	}

	return s, nil

}

func (s snake) View() string {

	display := ""

	display += fmt.Sprintf("Points Scored : %v\n", snek.points)

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
		for j := 0; j < WIDTH; j++ {

			switch area[i][j] {
			case 0:
				display += EMPTY

			case 3:
				display += SNAKE_BODY

			case 2:
				display += FOOD

			case 1:
				switch snek.dir {

				case UP:
					display += SNAKE_UP

				case DOWN:
					display += SNAKE_DOWN

				case LEFT:
					display += SNAKE_LEFT

				case RIGHT:
					display += SNAKE_RIGHT
				}

			}

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
