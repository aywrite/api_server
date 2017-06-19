package main

// Structs
type Index struct {
	Name string
}

type User struct {
	Name string `json:"name" validate:"min=4"`
	Age  int    `json:"age"  validate:"min=18, max=150"`
}

type Project struct {
	Name  string  `json:"name"  validate:"nonzero, min=4, max=150"`
	Price JSONInt `json:"price"  validate:"optmin=10"`
	User  User    `json:"user"  validate:"nonzero"`
}
