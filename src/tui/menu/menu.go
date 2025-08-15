package menu

type Menu interface {
	Main(screenWidth int) []string
}

type Section struct {
	LineStart int
}
