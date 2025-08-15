package menu

type Menu interface {
	Main(screenWidth int) []string
	Content() string
}

type Section struct {
	LineStart int
}
