package tui

import (
	"flak/src/config"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type screen struct {
	cursorX int
	cursorY int
	width   int
	height  int
	config  config.Config
}

func InitScreen(data config.Config) screen {
	return screen{
		config:  data,
		cursorX: 0,
		cursorY: 0,
		width:   0,
		height:  0,
	}
}

var (
	// Inverted style for character under cursor
	invertedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")). // Black foreground
		Background(lipgloss.Color("7"))  // White background
)

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
	var view string
	texts := []string{
		"oke", "not oke",
	}

	// Top Border
	view += "┌" + strings.Repeat("─", screen.width) + "┐\n"

	for yAxis := 0; yAxis < screen.height; yAxis++ {
		view += "│"
		var text []rune
		if yAxis < len(texts) {
			text = []rune(texts[yAxis])
		}
		for xAxis := 0; xAxis < screen.width; xAxis++ {
			if xAxis == screen.cursorX && yAxis == screen.cursorY {
				if xAxis < len(text) {
					view += invertedStyle.Render(string(text[xAxis]))
				} else {
					view += "█"
				}
			} else {
				if xAxis < len(text) {
					view += string(text[xAxis])
				} else {
					view += " "
				}
			}
		}
		view += "│\n"
	}

	// bottom Border
	view += "└" + strings.Repeat("─", screen.width) + "┘"
	view += "\n[.]Stop/Start All [Space]Start/Stop [,]Reload  [/]Root Folder [']Change Port [;]Open in Web [?]Help"
	return view
}
