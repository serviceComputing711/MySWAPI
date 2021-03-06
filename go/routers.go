/*
 * SWAPI
 *
 * This is a RESTful API SWAPI written in GO. You can find out more about SWAPI at [https://www.swapi.co](https://www.swapi.co). And you can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/). 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"Login",
		strings.ToUpper("Get"),
		"/login",
		Login,
	},

	Route{
		"AddPeople",
		strings.ToUpper("Post"),
		"/people/{peopleId}",
		AddPeople,
	},

	Route{
		"DeletePeople",
		strings.ToUpper("Delete"),
		"/people/{peopleId}",
		DeletePeople,
	},

	Route{
		"GetPeopleById",
		strings.ToUpper("Get"),
		"/people/{peopleId}",
		GetPeopleById,
	},

	Route{
		"GetPeopleByPage",
		strings.ToUpper("Get"),
		"/people",
		GetPeopleByPage,
	},

	Route{
		"GetAllAPI",
		strings.ToUpper("Get"),
		"/root",
		GetAllAPI,
	},
}
