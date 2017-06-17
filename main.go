package main

import (
	"log"
	"os"
	"text/template"
)

var routesTemplate *template.Template
var handlersTemplate *template.Template
var validationsTemplate *template.Template

func init() {
	routesTemplate = template.Must(
		template.ParseFiles("./templates/RoutesTemplate.go"),
	)
	handlersTemplate = template.Must(
		template.ParseFiles("./templates/HandlersTemplate.go"),
	)
	validationsTemplate = template.Must(
		template.ParseFiles("./templates/ValidationsTemplate.go"),
	)
}

func buildRoutes(r Routes) {
	f, err := os.Create("../apiServer/Routes.go")
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = routesTemplate.Execute(f, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildHandlers(r Routes) {
	f, err := os.Create("../apiServer/Handlers.go")
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = handlersTemplate.Execute(f, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildValidations(r Routes) {
	f, err := os.Create("../apiServer/Validations.go")
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = validationsTemplate.Execute(f, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	routes := ParseRoutes()
	buildRoutes(routes)
	buildHandlers(routes)
	buildValidations(routes)
}
