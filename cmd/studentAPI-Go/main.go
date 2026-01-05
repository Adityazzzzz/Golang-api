package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Adityazzzzz/studentAPI-Go/intenal/config"
)

func main() {
	fmt.Println("Goland_api")
/*
TODO:
	load config
	db setup
	routes
	server
*/

//config
	cfg := config.MustLoad()

//db

//routes
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request){
		res.Write([]byte("welcome to api"))
	})

//server
	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	err:=server.ListenAndServe()
	if err!=nil {
		log.Fatal("Failed to start server")
	}

	fmt.Println("server started")
}