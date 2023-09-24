package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{

		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{

		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		//make a chan to listen for term signals and wait for the signals
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		//shut down the server
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	//run the server in the main go routine
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()

	if err != nil {
		errorLog.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
