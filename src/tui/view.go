package tui

import (
	"flak/src/config"
	"flak/src/tui/menu"
	"fmt"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Screen struct {
	cursorX    int
	cursorY    int
	showCursor bool
	width      int
	height     int
	menu       string
	config     config.Config
	listMenu   []string
}

func InitScreen(data config.Config) Screen {
	return Screen{
		config:     data,
		cursorX:    0,
		cursorY:    0,
		width:      0,
		height:     0,
		showCursor: true, // Start visible
		menu:       "Applications",
		listMenu: []string{
			"Applications",
			"Settings",
			"Registry",
			"Database",
			"Cron",
			"Generator",
			"ApiClient",
			"Regex",
			"NetStat",
			"Note",
		},
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

func (screen Screen) Init() tea.Cmd {
	return tick()
}

func (screen Screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		screen.width = msg.Width - 2
		screen.height = msg.Height - 3
		screen.cursorX = 0
		screen.cursorY = 4
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
		default:
			{
				if key := msg.String(); len(key) == 1 && key >= "0" && key <= "9" {
					num, _ := strconv.Atoi(key)
					screen.menu = screen.listMenu[num]
				}
			}
		}
	case tickMsg:
		screen.showCursor = !screen.showCursor
		return screen, tick() // Schedule next tick
	}
	return screen, nil
}

func (screen Screen) handleBlinking(text, invertedText string) string {
	if screen.showCursor {
		return text
	}
	return invertedText
}
func (screen Screen) menuSection() string {
	menu := ""
	for key, value := range screen.listMenu {
		if screen.menu == value {
			menu += fmt.Sprintf("[*]%s ", value)
		} else {
			menu += fmt.Sprintf("[%d]%s ", key, value)
		}
	}
	paddingX := max(screen.width-len(menu), 0)
	view := strings.Repeat(" ", paddingX/2) + menu + strings.Repeat(" ", paddingX/2)
	return view
}

func (screen Screen) View() string {
	texts := []string{}
	texts = append(texts, screen.menuSection())
	texts = append(texts, strings.Repeat("─", screen.width))
	texts = append(texts, menu.Header(screen.width))
	texts = append(texts, strings.Repeat("─", screen.width))
	texts = append(texts, screen.menu)

	// Top Border
	var view string
	view += "┌" + strings.Repeat("─", screen.width) + "┐\n"

	for yAxis := 0; yAxis < screen.height; yAxis++ {
		view += "│"
		var text []rune
		if yAxis < len(texts) {
			text = []rune(texts[yAxis])
		}
		for xAxis := 0; xAxis < screen.width; xAxis++ {
			isCurrentCursor := xAxis == screen.cursorX && yAxis == screen.cursorY
			isCursorInFrontChar := xAxis < len(text)
			if isCurrentCursor && isCursorInFrontChar {
				view += screen.handleBlinking(string(text[xAxis]), invertedStyle.Render(string(text[xAxis])))
			} else if isCurrentCursor {
				view += screen.handleBlinking("█", " ")
			} else if isCursorInFrontChar {
				view += string(text[xAxis])
			} else {
				view += " "
			}
		}
		view += "│\n"
	}

	// bottom Border
	view += "└" + strings.Repeat("─", screen.width) + "┘"
	view += "\n[.]Stop/Start All [Space]Start/Stop [,]Reload  [/]Root Folder [']Change Port [;]Open in Web [?]Help"
	return view
}
