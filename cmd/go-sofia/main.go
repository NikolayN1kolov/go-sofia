package main

import (
	"os"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/NikolayN1kolov/go-sofia/internal/diagnostics"

)

func main() {
	log.Print("Starting the application...")

	blport := os.Getenv("PORT")
	if len(blport) == 0{
		log.Fatal("The application port should be set")
	}
	diagnosticsPort := os.Getenv("DIAG_PORT")
	if len(diagnosticsPort) == 0{
		log.Fatal("The diagnostics port should be set")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	go func() {
		log.Print("The applicaiton server is preaparing to handle connections..")
		server := http.Server{
			Addr: ":"+ blport,
			Handler: router,
		}
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Print("The diagnostics server is preaparing to handle connections..")
    diagnostics := diagnostics.NewDiagnostics()
	err := http.ListenAndServe(":"+ diagnosticsPort, diagnostics)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Print("The hello handler was called")
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
