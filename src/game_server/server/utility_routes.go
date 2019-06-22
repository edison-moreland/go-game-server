package server

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"net/http"
)

func AddUtilityRoutes(router *fasthttprouter.Router) {
	// TODO Switch to using sub routers

	router.GET("/ping", pingHandler)
}

// pingHandler is used by client to check that server is online
func pingHandler(ctx *fasthttp.RequestCtx) {
	// TODO: Return server information. players online, etc
	ctx.Response.SetStatusCode(http.StatusNoContent)
	return
}
