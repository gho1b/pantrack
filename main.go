package main

import (
	"github.com/gho1b/pantrack/handlers/services"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	srv := service.New(service.Name("pantrack"))
	_ = srv.Handle(new(services.LocationType))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
