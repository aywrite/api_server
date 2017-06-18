package main

type Config struct {
	Host string
	Port string
}

func ParseConfig() Config {
	return Config{
		"",
		"8080",
	}
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc string
	Params      interface{}
}

type Routes []Route

func ParseRoutes() Routes {
	var routes = Routes{
		Route{
			"Index",
			"Post",
			"/",
			"root",
			"Index",
		},
		Route{
			"Projects",
			"Post",
			"/projects/",
			"projectsGet",
			"Project",
		},
		Route{
			"Users",
			"Post",
			"/users/",
			"usersGet",
			"User",
		},
	}
	return routes
}
