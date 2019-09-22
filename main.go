package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"github.com/txvier/user-srv/api/proto/user"
	"github.com/txvier/user-srv/internal/service/endpoint"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.txvier.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Server().Init(
		server.Wait(nil),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(endpoint.UserHandler))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
