// Parsers for converting api schema and server config into go stucts.
// The resulting go structs are then used for server code generation.
package main

import "path/filepath"

type Config struct {
	Host     string // dns to run the server on
	Port     string // port to run the server on
	Location string // directory path where generated code should be created. Directory will be created if it does not exist.
}

// Parse the api server config file and return options as server config struct
// currently stubbed out to just return statically defined server config instead
func ParseConfig() Config {
	path, err := filepath.Abs("../api-server")
	if err != nil {
		panic(err)
	}
	return Config{
		"",
		"8080",
		path,
	}
}

// definition of each endpoint usually these would be more appropriate types
// but each value is only used for code generation so its easier to leave it as
// string (for now at least) although this does prevent any useful generation
// time validations
type Route struct {
	Name        string // name of the endpoint
	Method      string // method to use [GET, PUT, PATCH, POST, DELETE]
	Pattern     string // pattern to use to match the route
	HandlerFunc string // name to use for the handler function, must be a unique, valid go identifier
	Params      string // name of the struct type which defines the request parameters (currently must be defined in TypesTemplate.go)
}

type Routes []Route

// Parse the schema file and return the route structs for each endpoint
// currently stubbed out to just return statically defined routes instead
func ParseRoutes() Routes {
	return Routes{
		Route{
			"Index",
			"Post",
			"/",
			"root",
			"Index",
		}, // top level endpoint
		Route{
			"Projects",
			"Post",
			"/projects/",
			"projectsPost",
			"Project",
		}, // projects post endpoint
		Route{
			"Users",
			"Post",
			"/users/",
			"usersPost",
			"User",
		}, // users post endpoint
	}
}
