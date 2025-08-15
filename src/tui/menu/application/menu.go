package application

import (
	"flak/src/config"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Section struct {
	LineStart int
}

func New() Section {
	return Section{
		LineStart: 4,
	}
}

func (app Section) Main(screenWidth int) []string {
	headers := []string{
		"Name",
		"Port",
		"Status",
	}
	repeat := max(screenWidth-2, 0)
	sectionLength := repeat / len(headers)

	headerTop := ""
	headerContent := ""
	headerCenter := ""
	headerButton := ""
	for key, header := range headers {
		textContent := ""
		textTop := ""
		textCenter := ""
		textButton := ""
		if key != 0 {
			textContent = "│"
			textTop = "┬"
			textCenter = "┼"
			textButton = "┴"

		}
		if (sectionLength - len(header)) > 0 {
			textContent += " " + header
			textContent += strings.Repeat(" ", max(sectionLength-len(textContent), 0))
			headerContent += textContent
			headerCenter += textCenter + strings.Repeat("┈", utf8.RuneCountInString(textContent)-key)
			headerTop += textTop + strings.Repeat("─", utf8.RuneCountInString(textContent)-key)
			headerButton += textButton + strings.Repeat("─", utf8.RuneCountInString(textContent)-key)
		}
	}
	leftoverText := strings.Repeat(" ", max(repeat-utf8.RuneCountInString(headerContent), 0))
	leftoverTextDash := strings.Repeat("─", max(repeat-utf8.RuneCountInString(headerCenter), 0))
	texts := []string{}
	texts = append(texts, "╭"+headerTop+leftoverTextDash+"╮")
	texts = append(texts, "│"+headerContent+leftoverText+"│")
	texts = append(texts, "├"+headerCenter+leftoverTextDash+"┤")
	for _, s := range config.LoadConfig().Service {
		listService := ""
		listService += s.Name + strings.Repeat(" ", max(sectionLength-len(s.Name), 0))
		listService += "│" + strconv.Itoa(s.Port) + strings.Repeat(" ", max(sectionLength-utf8.RuneCountInString("│"+strconv.Itoa(s.Port))-2, 0))
		listService += "│ -" + strings.Repeat(" ", max(sectionLength-utf8.RuneCountInString("│ -"), 0)+3)
		texts = append(texts, "│"+listService+"│")
	}
	texts = append(texts, "├"+headerButton+leftoverTextDash+"┤")
	return texts
}

func (app Section) Content() []string {
	texts := []string{}
	texts = append(texts, "App Menu")
	return texts
}
