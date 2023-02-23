package main

import (
	log "github.com/sirupsen/logrus"
	"go-comments-rest-api/internal/comment"
	database "go-comments-rest-api/internal/db"
	transportHTTP "go-comments-rest-api/internal/transport/http"
)

// Run - responsible for the instantiation
// and startup of our application
// Run - sets up our application
// Run - sets up our application
func Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Setting Up Our APP")

	var err error
	store, err := database.NewDatabase()
	if err != nil {
		log.Error("failed to setup connection to the database")
		return err
	}
	err = store.MigrateDB()
	if err != nil {
		log.Error("failed to setup database")
		return err
	}

	commentService := comment.NewService(store)
	handler := transportHTTP.NewHandler(commentService)

	if err := handler.Serve(); err != nil {
		log.Error("failed to gracefully serve our application")
		return err
	}

	return nil
}
func main() {
	if err := Run(); err != nil {
		log.Error(err)
		log.Fatal("Error starting up our REST API")
	}
}
