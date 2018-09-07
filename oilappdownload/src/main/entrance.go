package main

import (
	"net/http"
	"myapp"
	"os"
	"os/signal"
	"log"
	"time"
)





/**
提供APP下载
*/
func main() {

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.HandleFunc("/downloadapp",myapp.APPdownload)


	server := http.Server{
		Addr:         ":9001",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under requeset!!")
		} else {
			log.Fatal("Server closed unexpecteed!!")
		}

	}

}
