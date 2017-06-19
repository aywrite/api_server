package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func ensureServerDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

func buildServer(c Config) {
	var serverTemplate *template.Template
	serverTemplate = template.Must(
		template.ParseFiles("./templates/ServerTemplate.go"),
	)
	f, err := os.Create(filepath.Join(c.Location, "/server.go"))
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = serverTemplate.Execute(f, c)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildTypes(r Routes, c Config) {
	var typesTemplate *template.Template
	typesTemplate = template.Must(
		template.ParseFiles("./templates/TypesTemplate.go"),
	)
	f, err := os.Create(filepath.Join(c.Location, "/types.go"))
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = typesTemplate.Execute(f, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildJSONTypes(r Routes, c Config) {
	var typesTemplate *template.Template
	typesTemplate = template.Must(
		template.ParseFiles("./templates/JSONTypesTemplate.go"),
	)
	f, err := os.Create(filepath.Join(c.Location, "/JSONtypes.go"))
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = typesTemplate.Execute(f, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildRoutes(r Routes, c Config) {
	var routesTemplate *template.Template
	routesTemplate = template.Must(
		template.ParseFiles("./templates/RoutesTemplate.go"),
	)
	f, err := os.Create(filepath.Join(c.Location, "/routes.go"))
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = routesTemplate.Execute(f, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildHandlers(r Routes, c Config) {
	var handlersTemplate *template.Template
	handlersTemplate = template.Must(
		template.ParseFiles("./templates/HandlersTemplate.go"),
	)
	f, err := os.Create(filepath.Join(c.Location, "/handlers.go"))
	if err != nil {
		log.Fatalln("create file: ", err)
	}
	err = handlersTemplate.Execute(f, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildValidations(r Routes, c Config) {
	var validationsTemplate *template.Template
	validationsTemplate = template.Must(
		template.ParseFiles("./templates/ValidationsTemplate.go"),
	)
	f, err := os.Create(filepath.Join(c.Location, "/validations.go"))
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
	ensureServerDirectory(config.Location)
	buildServer(config)
	buildTypes(routes, config)
	buildJSONTypes(routes, config)
	buildRoutes(routes, config)
	buildHandlers(routes, config)
	buildValidations(routes, config)
}
