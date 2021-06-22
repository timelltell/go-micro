package main

import (
	"bj38/handler"
	bj38 "bj38/proto/bj38"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	//"github.com/micro/go-micro/v2/registry"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	consulReg:=consul.NewRegistry()
	//consulReg:=consul.NewRegistry(func(options *registry.Options){
	//	options.Addrs=[]string{
	//		"127.0.0.1:8800",
	//	}
	//})
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.bj38"),
		micro.Registry(consulReg),
		micro.Version("latest"),
	)

	// Initialise service
	//service.Init()

	// Register Handler
	bj38.RegisterBj38Handler(service.Server(), new(handler.Bj38))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.service.bj38", service.Server(), new(subscriber.Bj38))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
