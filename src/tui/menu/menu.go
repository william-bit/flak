package menu

type Menu interface {
	Header(screenWidth int) string
	Content() string
}

type Section struct {
	LineStart int
}
