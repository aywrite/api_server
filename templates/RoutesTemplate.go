{{/* template used to generate routes */}}
// THIS FILE IS MACHINE GENERATED AND SHOULD
// NOT BE EDITED BY HAND

package main

import (
	"goji.io"
	"goji.io/pat"
)

func NewRouter() *goji.Mux {
	mux := goji.NewMux()
	{{range . -}}
	mux.HandleFunc(pat.{{.Method}}("{{.Pattern}}"), {{.HandlerFunc}})
	{{end -}}
	return mux
}
