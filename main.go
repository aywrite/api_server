package main

import (
	"log"
	"os"
	"text/template"
)

func buildServer(c Config) {
	var serverTemplate *template.Template
	serverTemplate = template.Must(
		template.ParseFiles("./templates/ServerTemplate.go"),
	)
	f, err := os.Create("../apiServer/server.go")
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = serverTemplate.Execute(f, c)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildTypes(r Routes) {
	var typesTemplate *template.Template
	typesTemplate = template.Must(
		template.ParseFiles("./templates/TypesTemplate.go"),
	)
	f, err := os.Create("../apiServer/types.go")
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = typesTemplate.Execute(f, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildRoutes(r Routes) {
	var routesTemplate *template.Template
	routesTemplate = template.Must(
		template.ParseFiles("./templates/RoutesTemplate.go"),
	)
	f, err := os.Create("../apiServer/routes.go")
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = routesTemplate.Execute(f, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildHandlers(r Routes) {
	var handlersTemplate *template.Template
	handlersTemplate = template.Must(
		template.ParseFiles("./templates/HandlersTemplate.go"),
	)
	f, err := os.Create("../apiServer/handlers.go")
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = handlersTemplate.Execute(f, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildValidations(r Routes) {
	var validationsTemplate *template.Template
	validationsTemplate = template.Must(
		template.ParseFiles("./templates/ValidationsTemplate.go"),
	)
	f, err := os.Create("../apiServer/validations.go")
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = validationsTemplate.Execute(f, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	config := ParseConfig()
	routes := ParseRoutes()
	buildServer(config)
	buildTypes(routes)
	buildRoutes(routes)
	buildHandlers(routes)
	buildValidations(routes)
}
