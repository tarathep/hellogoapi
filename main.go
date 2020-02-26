package main

import (
	"log"

	"github.com/tarathep/hellogoapi/kafka"
	"github.com/tarathep/hellogoapi/repository"

	"github.com/tarathep/hellogoapi/api"
	"github.com/tarathep/hellogoapi/route"
)

func main() {
	db, err := repository.Init("mongodb://admin:password@10.138.36.166:27021")
	if err != nil {
		log.Panic(err)
	}
	kafkaProducer := kafka.InitProducer("10.138.36.165:9092")

	go kafka.Consumer{
		Broker:   "10.138.36.165:9092",
		Version:  "2.1.1",
		Group:    "PHX-ORDER-CREATED-GROUP1",
		Topic:    "PHX-ORDER-CREATED",
		Assignor: "range",
		Oldest:   true,
		Verbose:  false,
	}.Run()

	route := route.Route{api.HelloHandler{db, kafkaProducer}}
	r := route.SetupRouter()
	r.Run()

}
