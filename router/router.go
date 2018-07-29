package router

import (
	"github.com/julienschmidt/httprouter"

	"github.com/bc-class/controller"
)

func GenRouter() *httprouter.Router {
	router := httprouter.New()

	// list all service
	router.GET("/list/deployment", controller.ListDeployment)

	// list all pods for deployment
	router.GET("/list/pod/:deployment", controller.ListPods)

	// scale deployment
	router.GET("/scale/:deployment/:amount", controller.ScaleDeployment)

	// delete pod
	router.DELETE("/pod/:name", controller.DeletePod)

	// create deployment
	router.POST("/create/deployment", controller.CreateDeployment)

	return router

}
