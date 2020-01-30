package main

import (
	"log"

	"github.com/tarathep/hellogoapi/repository"

	"github.com/tarathep/hellogoapi/api"
	"github.com/tarathep/hellogoapi/route"
)

func main() {
	db, err := repository.Init("mongodb://admin:password@localhost:27017")
	if err != nil {
		log.Panic(err)
	}

	route := route.Route{api.HelloHandler{db}}
	r := route.SetupRouter()
	r.Run()

}
