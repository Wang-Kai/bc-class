package router

import (
	"github.com/julienschmidt/httprouter"

	"github.com/bc-class/controller"
)

func GenRouter() *httprouter.Router {
	router := httprouter.New()

	// list all service
	router.GET("/list/deployment", controller.ListDeployment)
	// scale a service
	router.GET("/access/:deployment/:user", controller.HandleAccess)

	return router
}
