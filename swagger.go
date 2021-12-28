package main

import (
	"net/http"

	"github.com/davidebianchi/gswagger/apirouter"
	"github.com/labstack/echo/v4"
)

func newSwaggerEcho(router *echo.Echo) apirouter.Router {
	return echoRouter{router: router}
}

type echoRouter struct {
	router *echo.Echo
}

func (r echoRouter) AddRoute(path, method string, handler apirouter.HandlerFunc) apirouter.Route {
	return r.router.Add(method, path, echo.WrapHandler(http.HandlerFunc(handler)))
}
