package tui

import (
	"flak/src/config"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type screen struct {
	cursorX    int
	cursorY    int
	showCursor bool
	width      int
	height     int
	config     config.Config
}

func InitScreen(data config.Config) screen {
	return screen{
		config:     data,
		cursorX:    0,
		cursorY:    0,
		width:      0,
		height:     0,
		showCursor: true, // Start visible
	}
}

var (
	// Inverted style for character under cursor
	invertedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")). // Black foreground
		Background(lipgloss.Color("7"))  // White background
)

// Message to trigger cursor blink
type tickMsg time.Time

// Every 500ms, send a tickMsg to toggle cursor visibility
func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (screen screen) Init() tea.Cmd {
	return tick()
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
	case tickMsg:
		screen.showCursor = !screen.showCursor
		return screen, tick() // Schedule next tick
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
					if screen.showCursor {
						view += invertedStyle.Render(string(text[xAxis]))
					} else {
						view += string(text[xAxis])
					}
				} else {
					if screen.showCursor {
						view += "█"
					} else {
						view += " "
					}
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
