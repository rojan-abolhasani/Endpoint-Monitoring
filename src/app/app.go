package app

import (
	"monitor/config"
	"monitor/router"
	"net/http"
)

func Run() {
	mux := router.Router()
	// intializing the server
	server := &http.Server{
		Addr:         config.Addr,
		IdleTimeout:  config.IdleTimeOut,
		ReadTimeout:  config.ReadTimeOut,
		WriteTimeout: config.WriteTimeOut,
		Handler:      mux,
	}
	// create a new go routine to monitor the links
	go monitor()
	// run the server
	server.ListenAndServe()
}
