package main

import (
	"bytes"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/prefork"
)

func main() {
	methodGet := []byte("GET")
	path := []byte("/")
	handler := func(ctx *fasthttp.RequestCtx) {
		if bytes.Equal(ctx.Method(), methodGet) && bytes.Equal(ctx.Path(), path) {
			ctx.Response.SetBodyString("Hello, 世界!")
		}
	}

	server := &fasthttp.Server{
		Handler:                       handler,
		DisableHeaderNamesNormalizing: true,
	}

	if err := prefork.New(server).ListenAndServe(":3000"); err != nil {
		panic(err)
	}
}
