package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/dukinm/go-musthave-shortener/internal/app"
)

func main() {
	app := app.NewRouter().Run(":8080")

	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer done()

	<-ctx.Done()
	log.Println("Service down.")
	app.Stop(ctx)
}
