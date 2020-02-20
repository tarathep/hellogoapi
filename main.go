package main

import (
	"log"

	"github.com/tarathep/hellogoapi/event"

	"github.com/tarathep/hellogoapi/repository"

	"github.com/tarathep/hellogoapi/api"
	"github.com/tarathep/hellogoapi/route"
)

func main() {
	db, err := repository.Init("mongodb://admin:password@10.138.36.166:27021")
	if err != nil {
		log.Panic(err)
	}
	kafka := event.Init("10.138.36.165:9092")

	route := route.Route{api.HelloHandler{db, kafka}}
	r := route.SetupRouter()
	r.Run()

}
