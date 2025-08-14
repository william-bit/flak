package generator

import "strings"

type Section struct {
	LineStart int
}

func New() Section {
	return Section{
		LineStart: 4,
	}
}

func (app Section) Header(screenWidth int) string {
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
	return s
}

func (app Section) Content() string {
	return "Generator Menu"
}
