package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type screen struct {
	cursorX int
	cursorY int
	width   int
	height  int
}

func InitScreen() screen {
	return screen{
		cursorX: 0,
		cursorY: 0,
		width:   0,
		height:  0,
	}
}

func (screen screen) Init() tea.Cmd {
	return nil
}

func (screen screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		screen.width = msg.Width - 2
		screen.height = msg.Height - 3
		screen.cursorX = 0
		screen.cursorY = 0
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return screen, tea.Quit
		case "up", "k":
			if screen.cursorY > 0 {
				screen.cursorY--
			}
		case "down", "j":
			if screen.cursorY < screen.height-1 {
				screen.cursorY++
			}
		case "left", "h":
			if screen.cursorX > 0 {
				screen.cursorX--
			}
		case "right", "l":
			if screen.cursorX < screen.width-1 {
				screen.cursorX++
			}
		}
	}
	return screen, nil
}

func (screen screen) View() string {
	var s string

	// Top Border
	s += "╭" + strings.Repeat("─", screen.width) + "╮\n"

	for yAxis := 0; yAxis < screen.height; yAxis++ {
		s += "│"
		for xAxis := 0; xAxis < screen.width; xAxis++ {
			if xAxis == screen.cursorX && yAxis == screen.cursorY {
				s += "█"
			} else {
				s += " "
			}
		}
		s += "│\n"
	}

	// bottom Border
	s += "╰" + strings.Repeat("─", screen.width) + "╯"
	s += "\n[.]Stop/Start All [Space]Start/Stop [,]Reload  [/]Root Folder [']Change Port [;]Open in Web [?]Help"
	return s
}
