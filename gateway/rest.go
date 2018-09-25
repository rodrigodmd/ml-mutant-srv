package main

import (
	"log"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-web"
	"github.com/rodrigodmd/ml-mutant-srv/gateway/handler"
)

func main() {
	// Create API service
	service := web.NewService(
		web.Name("go.micro.api.mutant"),
	)

	service.Init()

	wc := restful.NewContainer()
	// Create RESTful handler
	say := handler.Dna{}

	ws1 := restful.WebService{}
	ws1.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws1.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws1.Path("/mutant")
	ws1.Route(ws1.POST("/").To(say.IsMutant))

	ws2 := restful.WebService{}
	ws2.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws2.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws2.Path("/stat")
	ws2.Route(ws1.POST("/").To(say.Stat))

	wc.Add(&ws1)
	wc.Add(&ws2)

	// Register Handler
	service.Handle("/", wc)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
