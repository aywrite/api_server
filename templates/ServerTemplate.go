{{/* Server template */}}
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var (
	ErrDecodingError = errors.New("decoding error")
)

type EndpointParams interface {
	Validate(r *http.Request) error
}

func decodeAndValidate(r *http.Request, params EndpointParams) error {
	if err := json.NewDecoder(r.Body).Decode(params); err != nil {
		fmt.Println(err.Error())
		fmt.Println(err)
		return ErrDecodingError
	}
	defer r.Body.Close()
	return params.Validate(r)
}

func jsonResponse(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(v); err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := NewRouter()
	http.ListenAndServe("{{.Host}}:{{.Port}}", router)
}
