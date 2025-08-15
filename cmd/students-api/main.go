package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mubinkg/student-management/internal/config"
)

func main() {
	fmt.Println("Hello students")
	config := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Weelcome to students api"))
	})

	server := http.Server{
		Addr:    config.Address,
		Handler: router,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, syscall.SIGINT, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Printf("Server started: http://%s", config.Address)

		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Info("Fialed to shutdown the server")
	}

	slog.Info("Server shutdown successfully")
}
