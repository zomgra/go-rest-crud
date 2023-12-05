package entities

type Book struct {
	Id       string  `json:id`
	Title    string  `json:title`
	Author   *Author `json:author`
	AuthorId int     `json:author`
	Isnbn    string  `json:isnbn`
}
