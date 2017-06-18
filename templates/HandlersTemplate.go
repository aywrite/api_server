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
		jsonResponse(w, http.StatusBadRequest, err)
    } else {
		jsonResponse(w, http.StatusOK, "{{.Name}} ok")
	}
}
{{end}}
