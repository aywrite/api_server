{{/* template used to generate request validators */}}
// THIS FILE IS MACHINE GENERATED AND SHOULD
// NOT BE EDITED BY HAND

package main

import (
	"net/http"
)

{{range .}}
//{{.Name}} Validations for {{.Params}}
func (params {{.Params}}) Validate(r *http.Request) error {
    return nil
}
{{end}}
