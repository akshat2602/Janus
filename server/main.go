package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/akshat2602/Janus/cmd/webserver"
	"github.com/akshat2602/Janus/utils/logger"
	"github.com/rs/cors"
)

func main() {
	logger.InitializeLogger()

	// Initialize the web server
	webServer := webserver.InitializeWebServer()

	corsHandler := cors.New(
		cors.Options{
			// TODO: Change this to the frontend URL
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
		},
	).Handler(webServer)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      corsHandler,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	logger.Logger.Info("Starting the server on port 8080")
	err := s.ListenAndServe()
	if err != nil {
		logger.Logger.Sugar().Error("Error starting the server: ", err)
	}

	fmt.Println("Hello, World!")
}
