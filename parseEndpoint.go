package main

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc string
	Params      interface{}
}

type Index struct {
	Name string
}

type User struct {
	Name string
	Age  int
}

type Project struct {
	Name  string
	Price float32
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
