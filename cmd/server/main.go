package main

import (
	"context"
	"fmt"
	database "go-comments-rest-api/db"
)

// Run - responsible for the instantiation
// and startup of our application
func Run() error {
	fmt.Println("starting our application")
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
