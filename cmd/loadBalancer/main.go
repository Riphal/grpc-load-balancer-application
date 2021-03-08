package main

import (
	"log"

	core "github.com/Riphal/grpc-load-balancer-application"
	pkgerrors "github.com/pkg/errors"
)

func main() {
	app := core.NewApp()

	registerRoutes(app)

	if err := app.Server.ListenAndServe(); err != nil {
		log.Fatalln(pkgerrors.Wrap(err, "failed to run server"))
	}
}
