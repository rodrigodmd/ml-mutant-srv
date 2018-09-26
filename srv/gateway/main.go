package main

import (
	"log"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-web"
	"github.com/rodrigodmd/ml-mutant-srv/srv/gateway/handler"
)

func main() {
	// Create API service
	service := web.NewService(
		web.Name("go.micro.api.api"),
	)

	service.Init()

	wc := restful.NewContainer()
	// Create RESTful handler
	dna := new(handler.Dna)

	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON, restful.MIME_XML)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/api")
	ws.Route(ws.POST("/mutant").To(dna.Mutant))
	ws.Route(ws.GET("/stats").To(dna.Stat))

	wc.Add(ws)

	// Register Handler
	service.Handle("/", wc)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
