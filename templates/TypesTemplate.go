package main

import (
	"errors"
)

// Structs
type Index struct {
	Name string
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Project struct {
	Name  string
	Price float32
}

// Custom Errors
var (
	ErrInvalidName = errors.New("invalid name")
)
