package entities

type Book struct {
	Id     string  `json:id`
	Title  string  `json:title`
	Author *Author `json:author`
	Isnbn  string  `json:isnbn`
}
