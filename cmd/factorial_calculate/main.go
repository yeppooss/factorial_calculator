package main

import (
	"factorial_calculate/internal/config"
	"factorial_calculate/internal/http-server/handlers"
	"factorial_calculate/internal/http-server/middlewares"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	cfg := config.MustLoad()

	router := httprouter.New()
	router.Handle("POST", "/calculate", middlewares.CheckParamsMiddleware(handlers.FactorialHandler))

	log.Fatalln(http.ListenAndServe(cfg.Address, router))
}
