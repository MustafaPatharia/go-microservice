package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/MustafaPatharia/go-microservice/application"
)

func main() {
	app := application.New() // app is a pointer to App (memory address)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Printf("Error starting application: %v\n", err)
	}

}
