package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/madhuri/pet-client/handlers"
	"github.com/madhuri/pet-client/service"
	"github.com/sirupsen/logrus"
)

func main() {
	baseURL, ok := os.LookupEnv("PET_BASE_URL")
	if !ok {
		logrus.Error("PET_BASE_URL env variable is not set")
		return
	}
	petClient := service.NewPetStore(baseURL)
	// Capture Ctrl-C
	ctx := context.Background()
	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	//---------------------------------------------------
	router := mux.NewRouter()
	handler := handlers.NewPetHandel(petClient)
	fmt.Println("started listening on port : ", 8080)
	router.Handle("/pets", http.HandlerFunc(handler.CreatePet)).Methods("POST")
	router.Handle("/pets/{id}", http.HandlerFunc(handler.GetPet)).Methods("GET")
	router.Handle("/pets", http.HandlerFunc(handler.GetPets)).Methods("GET")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("status", "fatal", "err", err)
	}

	//-------------capturing the ctrl + c event----------------------
	select {
	case <-c:
		log.Println("cancel operation")
		cancel()

	case <-ctx.Done():
		time.Sleep(600 * time.Millisecond)
	}

	log.Println("done")
}
