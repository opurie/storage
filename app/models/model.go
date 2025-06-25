package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Item struct {
	ID int
	Name string
	Category Category
	Location Location
	Owner User
	Description string
	Tags []Tag
}

type Category struct {
	ID int
	Name string
}
type Location struct {
	ID int
	Name string
}

type Tag struct{
	Name string
}