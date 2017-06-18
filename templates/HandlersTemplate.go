{{/* template used to handler */}}
// THIS FILE IS MACHINE GENERATED AND SHOULD
// NOT BE EDITED BY HAND

package main

import (
	"net/http"
)

func jsonError(w http.ResponseWriter, e error) error {
	switch e.(type) {
	case Failure:
		f := e.(Failure)
		return f.jsonResponse(w)
	default:
		return Failure{
			code: http.StatusInternalServerError,
			ErrorCode: "UnknownError",
			Message: "unknown error occured",
		}.jsonResponse(w)
	}
}

{{range .}}
// Handler for {{.Name}} [{{.Method}} {{.Pattern}}]
func {{.HandlerFunc}}(w http.ResponseWriter, r *http.Request) {
    var params *{{.Params}} = &{{.Params}}{}
    err := decodeAndValidate(r, params)
	if err != nil {
		_ = jsonError(w, err)
    } else {
		s := Success{
			code: http.StatusOK,
			Result: "{{.Name}} ok",
		}
		_ = s.jsonResponse(w)
	}
}
{{end}}
