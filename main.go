package main

import (
	"github.com/Dadard29/go-api-utils/API"
	"github.com/Dadard29/go-api-utils/database"
	"github.com/Dadard29/go-api-utils/service"
	"github.com/Dadard29/go-subscription-connector/subChecker"
	"github.com/Dadard29/go-todo/api"
	"github.com/Dadard29/go-todo/controllers"
	"net/http"
)

var routes = service.RouteMapping{
	"/one-time/task": service.Route{
		Description: "manage one-time task",
		MethodMapping: service.MethodMapping{
			http.MethodGet:    controllers.OneTimeGet,
			http.MethodPost:   controllers.OneTimePost,
			http.MethodPut:    controllers.OneTimePut,
			http.MethodDelete: controllers.OneTimeDelete,
		},
	},
	"/one-time/task/list": service.Route{
		Description:   "manage task list",
		MethodMapping: service.MethodMapping{
			http.MethodGet: controllers.OneTimeListGet,
		},
	},
	"/periodic/task": service.Route{
		Description: "manage periodic task",
		MethodMapping: service.MethodMapping{
			http.MethodGet:    controllers.PeriodicGet,
			http.MethodPost:   controllers.PeriodicPost,
			http.MethodPut:    controllers.PeriodicPut,
			http.MethodDelete: controllers.PeriodicDelete,
		},
	},
}

func main() {
	var err error
	api.Api = API.NewAPI("Todo", "config/config.json",
		routes, true)

	dbConfig, err := api.Api.Config.GetSubcategoryFromFile("api", "db")
	api.Api.Logger.CheckErrFatal(err)
	api.Api.Database = database.NewConnector(dbConfig, true, []interface{}{})

	controllers.Sc = subChecker.NewSubChecker(api.Api.Config.GetEnv("HOST_SUB"))

	api.Api.Service.Start()

	api.Api.Service.Stop()
}
