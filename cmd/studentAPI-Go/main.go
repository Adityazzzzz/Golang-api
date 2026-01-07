package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	


	//----------------------we prefer to run server in go-routines---------------------------

	checkisInterrupt := make(chan os.Signal,1)
	signal.Notify(checkisInterrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) //as a notification if there is interrupt

	go func(){
		err:=server.ListenAndServe()
		if err!=nil {
			log.Fatal("Failed to start server")
		}
	}()
	
	<-checkisInterrupt




	// structured logs
	slog.Info("shutting down the server")

	// we cant directly shutdown. so,we use timer
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}