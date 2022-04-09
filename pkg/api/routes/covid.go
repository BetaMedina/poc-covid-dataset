package routes

import (
	"dataset-covid/pkg/api/controllers"
)

var CovidRoutes = []Routes{
	{
		Uri:         "/covid/region/{state}",
		Method:      "GET",
		Func:        controllers.FindByState,
		RequestAuth: false,
	},
	{
		Uri:         "/covid/city/{city}",
		Method:      "GET",
		Func:        controllers.FindByCity,
		RequestAuth: false,
	},
}
