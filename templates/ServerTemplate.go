{{/* Server template */}}
package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"./mapstructure"
)

var (
	ErrDecodingError = errors.New("decoding error")
)

type EndpointParams interface {
	Validate(r *http.Request) error
}

type ApiResponse interface {
	jsonResponse(http.ResponseWriter) error
}

type Success struct {
	Status string // success
	code int // http status code
	Result interface{}// result object or message
}

type Failure struct {
	Status string // error
	code int // http response
	ErrorCode string // error category
	Message interface{}// error messages
}

func (e Failure) Error() string { return e.ErrorCode }

func (s Success) jsonResponse(w http.ResponseWriter) error {
	s.Status = "success"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(s.code)
	encoder := json.NewEncoder(w)
	return encoder.Encode(s)
}

func (f Failure) jsonResponse(w http.ResponseWriter) error {
	f.Status = "error"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(f.code)
	encoder := json.NewEncoder(w)
	return encoder.Encode(f)
}

// return a struct which has some kind of json method in it
// should be two structs, one for errors, one for successes
func decodeAndValidate(r *http.Request, params EndpointParams) error {
	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Print(err)
		f := Failure{
			code: http.StatusBadRequest,
			ErrorCode: "BadJson",
			Message: "Request json missing or malformed",
		}
		return f
	}
	if err := mapstructure.Decode(body, params); err != nil {
		log.Print(err)
		f := Failure{
			code: http.StatusBadRequest,
			ErrorCode: "InvalidJson",
			Message: err,
		}
		return f
	}
	defer r.Body.Close()
	if err := params.Validate(r); err != nil {
		log.Print(err)
		f := Failure{
			code: http.StatusBadRequest,
			ErrorCode: "InvalidJson",
			Message: err,
		}
		return f
	}
	return nil
}


func main() {
	router := NewRouter()
	http.ListenAndServe("{{.Host}}:{{.Port}}", router)
}
