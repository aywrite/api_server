{{/* template used to handler */}}
// THIS FILE IS MACHINE GENERATED AND SHOULD
// NOT BE EDITED BY HAND

package main

import (
	"net/http"
	"encoding/json"
)

{{range .}}
//{{.Name}} Handler for [{{.Method}}] {{.Pattern}}
func {{.HandlerFunc}}(w http.ResponseWriter, r *http.Request) {
    // we have an input structure, of type Thing
    var params *{{.Params}} = &{{.Params}}{}
    // we want to decode and validate Thing from request body
    err := decodeAndValidate(r, params)
    // there was an error with Thing
    if err != nil {
        // send a bad request back to the caller
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
        return
    }
    // it was decoded, and validated properly, success
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("{{.Name}} ok"))
}
{{end}}
