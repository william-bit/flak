package netstat

import "strings"

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
	sectionLength := screenWidth / len(headers)

	s := ""
	for key, header := range headers {
		if (sectionLength - len(header)) > 0 {
			s += " " + header + strings.Repeat(" ", sectionLength-len(" "+header))
			if key != len(headers)-1 {
				s += "│"
			}
		}
	}
	repeat := max(screenWidth-2, 0)
	texts := []string{}
	texts = append(texts, "┌"+strings.Repeat("─", repeat)+"┐")
	content := app.Content()
	texts = append(texts, "│"+content+strings.Repeat("─", repeat-len(content))+"│")
	return texts
}

func (app Section) Content() string {
	return "App Menu"
}
