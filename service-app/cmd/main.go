package main

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-app/auth"
	"service-app/core/inventory"
	"service-app/core/user"
	"service-app/database"
	"service-app/handlers"
	"time"
)

// go mod tidy // run it first time to set up this project and its deps,
func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	logging, err := setupLogger()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = startApp(logging)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func startApp(log *zap.Logger) error {
	// =========================================================================
	// Start Database
	log.Info("main : Started : Initializing db support")
	db, err := database.Open()
	if err != nil {
		return fmt.Errorf("connecting to db %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := db.Ping(ctx); err != nil {
		return err
	}

	// =========================================================================
	// Initialize Service layer support

	us, err := user.NewService(db)
	if err != nil {
		return err
	}
	inv, err := inventory.NewService(db)
	if err != nil {
		return err
	}

	// =========================================================================
	// Initialize authentication support
	log.Info("main : Started : Initializing authentication support")
	privatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		return fmt.Errorf("reading auth private key %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("parsing auth private key %w", err)
	}

	publicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		return fmt.Errorf("reading auth public key %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return fmt.Errorf("parsing auth public key %w", err)
	}

	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("constructing auth %w", err)
	}

	// =========================================================================

	// =========================================================================
	// Initialize service
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		Handler:      handlers.API(log, a, us, inv),
	}
	// channel to store any errors while setting up the service
	serverErrors := make(chan error, 1)
	go func() {
		//log.Printf("main: API listening on %s", api.Addr)
		log.Info("main: API listening on", zap.String("address", api.Addr))
		serverErrors <- api.ListenAndServe()
	}()

	//shutdown channel intercepts ctrl+c signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt) // this will notify the shutdown chan if someone presses ctr+c

	select {
	//this case would exec in case server is not able to start
	case err := <-serverErrors:
		return fmt.Errorf("server error %w", err)

	//this case runs when someone pressed ctrl+c
	case sig := <-shutdown:
		//log.Printf("main: %v : Start graceful shutdown", sig)
		log.Info("main: Start shutdown", zap.Any("signal", sig))
		ctx := context.Background()
		//creating a timeout of 10 seconds for our service to close the connections
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel() // clean the resources taken up by the context

		//Shutdown gracefully shuts down the server without interrupting any active connections.
		//Shutdown works by first closing all open listeners, then closing all idle connections,
		//and then waiting indefinitely for connections to return to idle and then shut down.
		err := api.Shutdown(ctx)
		if err != nil {
			//Close immediately closes all active net.Listeners
			err = api.Close() // forcing shutdown
			return fmt.Errorf("could not stop server gracefully %w", err)
		}
	}

	// =========================================================================
	return nil

}
