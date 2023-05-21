package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/ahmadaidin/ginplating/config"
	"github.com/ahmadaidin/ginplating/controller/httpctrl"
	"github.com/ahmadaidin/ginplating/controller/httpctrl/bookctrl"
	"github.com/ahmadaidin/ginplating/domain/repository"
	"github.com/ahmadaidin/ginplating/infrastructure/database"
)

func main() {
	configLoader, err := config.NewLoaderAndLoad()
	if err != nil {
		log.Fatal(err)
	}

	cfg := configLoader.Config()

	mongoDb := database.NewMongoDatabase(cfg.DatabaseURI, 10)

	bookRepo := repository.NewBookRepository(mongoDb)
	bookCtrl := bookctrl.NewBookController(
		&configLoader,
		bookRepo,
	)
	httpHandler := httpctrl.NewGinHttpHandler(*bookCtrl)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: httpHandler.Engine,
	}

	configLoader.OnDatabaseURIChange(func(oldURI string, newURI string) {
		log.Println("calling mongodb.NewClient with new URI")
		mongoDb.NewClient(newURI, 10)
	})

	configLoader.OnPortChange(func(oldPort string, newPort string) {
		log.Printf("port changed from %s to %s. doing nothing to the server\n", oldPort, newPort)
	})

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
