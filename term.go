package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func tick() tea.Cmd {

	return tea.Tick(snek.duration, func(t time.Time) tea.Msg {
		return t
	})

}

func (s snake) Init() tea.Cmd {
	return tick()
}

func (s snake) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s snake) View() string {

	display := ""

	return display
}