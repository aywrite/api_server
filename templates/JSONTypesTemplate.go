package main

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"log"
)

type JSONInt struct {
	Value int  // value of the variable
	Valid bool // was the null passed as the value
	Set   bool // was the value passed at all
}

func (v *JSONInt) passedAsNull() bool {
	return v.Valid
}

func (v *JSONInt) passed() bool {
	return v.Set
}

func (v *JSONInt) value() int {
	return v.Value
}

type JSONVar interface {
	value() int
	passedAsNull() bool
	passed() bool
}

func (i *JSONInt) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp int
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

// Very simple validator with parameter
func optMin(v interface{}, param string) error {
	switch jsonV := v.(type) {
	case JSONInt:
		if jsonV.passed() {
			validation := "min=" + param
			return validator.Valid(jsonV.value(), validation)
		}
	default:
		log.Print("went wrong here")
		return validator.ErrUnsupported
	}
	return nil
}

func registerCustomValidators() {
	validator.SetValidationFunc("optmin", optMin)
}
