package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/I1820/wf/actions"
)

func main() {
	fmt.Println("18.20 at Sep 07 2016 7:20 IR721")

	srv := &http.Server{
		Addr:    ":6976",
		Handler: actions.App(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("API Service failed with %s", err)

		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("18.20 As always ... left me alone")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("API Service failed on exit: %s", err)

	}
}
