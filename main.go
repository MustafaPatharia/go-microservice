package main

import (
	"context"
	"fmt"

	"github.com/MustafaPatharia/go-microservice/application"
)

func main() {
	app := application.New() // app is a pointer to App (memory address)
	err := app.Start(context.TODO())

	if err != nil {
		fmt.Printf("Error starting application: %v\n", err)
	}

}
