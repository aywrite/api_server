{{/* Server template */}}
package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"fmt"
	"gopkg.in/validator.v2"
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

func jsonError(w http.ResponseWriter, e error) error {
	f, ok := e.(Failure)
	if !ok {
		return Failure{
			code: http.StatusInternalServerError,
			ErrorCode: "UnknownError",
			Message: "unknown error occured",
		}.jsonResponse(w)
	}
	return f.jsonResponse(w)
}

func (f Failure) jsonResponse(w http.ResponseWriter) error {
	f.Status = "error"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(f.code)
	encoder := json.NewEncoder(w)
	return encoder.Encode(f)
}

func decodingError(err error) error {
		log.Print(err)
		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			return Failure{
				code: http.StatusBadRequest,
				ErrorCode: "BadJson",
				Message: fmt.Sprintf("invalid type for parameter '%s': expected type '%v' got '%s'", e.Field, e.Type, e.Value),
			}
		case *json.SyntaxError:
			return Failure{
				code: http.StatusBadRequest,
				ErrorCode: "BadJson",
				Message: "Request json missing or malformed",
			}
		default:
			return err
		}
}

func decodeAndValidate(r *http.Request, params EndpointParams) error {
	if err := json.NewDecoder(r.Body).Decode(params); err != nil {
		return decodingError(err)
	}
	defer r.Body.Close()
	if errs := validator.Validate(params); errs != nil {
		log.Print(errs)
		return Failure{
			code: http.StatusBadRequest,
			ErrorCode: "InvalidJson",
			Message: errs,
		}
	}
	return nil
}

func main() {
	registerCustomValidators()
	router := NewRouter()
	http.ListenAndServe("{{.Host}}:{{.Port}}", router)
}
