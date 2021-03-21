package model

type Test struct {
	A string
	B string
}

type Note struct {
	Name  string
	Steps []Step
}

type Step struct {
	Title string
	Txt   string
	Url   *string
}
