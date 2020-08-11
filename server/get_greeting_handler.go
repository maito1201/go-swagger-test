package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/maitoprivate/go-swagger-test/server/gen/restapi/factory"
)

func GetGreeting(p factory.GetGreetingParams) middleware.Responder {
	payload := "hello " + *p.Name
	return factory.NewGetGreetingOK().WithPayload(payload)
}
