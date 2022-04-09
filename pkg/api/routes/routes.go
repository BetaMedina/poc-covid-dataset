package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	Uri         string
	Method      string
	Func        func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}

func config(r *mux.Router) *mux.Router {
	for _, route := range CovidRoutes {
		r.HandleFunc(route.Uri, route.Func).Methods(route.Method)
	}
	return r
}

func GenerateMuxRouter() *mux.Router {
	r := mux.NewRouter()
	return config(r)
}
