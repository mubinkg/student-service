package main

import (
	"fmt"
	"log"
	"net/http"

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

	fmt.Printf("Server started: http://%s", config.Address)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
