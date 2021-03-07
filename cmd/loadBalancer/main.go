package main

import (
	core "github.com/Riphal/grpc-load-balancer-application"
	pkgerrors "github.com/pkg/errors"
	goLog "log"
)

func main() {
	app := core.NewApp()

	registerRoutes(app)

	if err := app.Server.ListenAndServe(); err != nil {
		goLog.Fatalln(pkgerrors.Wrap(err, "failed to run server"))
	}
}
