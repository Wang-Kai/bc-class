package main

import (
	"github.com/urfave/negroni"

	"github.com/bc-class/middleware"
	"github.com/bc-class/router"
)

func main() {
	n := negroni.New()

	// catch panic
	n.Use(negroni.NewRecovery())

	n.Use(middleware.Logger())

	// bind router
	r := router.GenRouter()
	n.UseHandler(r)

	n.Run(":8848")
}
