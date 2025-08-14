package menu

import "strings"

func Header(screenWidth int) string {
	headers := []string{
		"Name",
		"Port",
		"Status",
	}
	sectionLength := screenWidth / len(headers)

	s := ""
	for key, header := range headers {
		if (sectionLength - len(header)) > 0 {
			s += header + strings.Repeat(" ", sectionLength-len(header))
			if key != len(headers)-1 {
				s += "│"
			}
		}
	}
	return s
}
