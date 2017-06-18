{{/* template used to handler */}}
// THIS FILE IS MACHINE GENERATED AND SHOULD
// NOT BE EDITED BY HAND

package main

import (
	"net/http"
)


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
