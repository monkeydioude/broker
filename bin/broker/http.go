package main

import (
	"log"

	"github.com/monkeydioude/moon"
	"github.com/monkeydioude/tools"
)

func healthcheck(r *moon.Request) ([]byte, int, error) {
	return tools.Response200([]byte("ok"))
}

func httpServer() {
	handler := moon.NewHandler()
	handler.AddHeader("Access-Control-Allow-Origin", "*")
	handler.MakeRouter(
		moon.Get("healthcheck$", healthcheck),
	)

	if err := moon.ServerRun(httpPort, handler); err != nil {
		log.Printf("[ERR ] Server crashed. Reason: %s", err)
	}
}
