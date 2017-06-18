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
	Price float64 `json:"price"  validate:"min=0"`
	User  User    `json:"user"  validate:"nonzero"`
}
