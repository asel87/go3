package main

import (
	"fmt"
	"log"
	"net/http"

	"assik3/pkg/store/postgres"
	delivery "assik3/services/contact/internal/delivery/http"
	repository "assik3/services/contact/internal/repository/storage/postgres"
	contactUseCase "assik3/services/contact/internal/useCase/contact"
	groupUseCase "assik3/services/contact/internal/useCase/group"
)

func main() {
	config := postgres.NewDBConfig("localhost", 5433, "postgres", "asel3127", "postgres")
	db, err := postgres.ConnectToDB(config)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	r := repository.New(db)
	ucGroup := groupUseCase.New(r)
	ucContact := contactUseCase.New(r)

	d := delivery.New(ucContact, ucGroup)

	addr := 4000
	addrStr := fmt.Sprintf(":%d", addr)

	log.Printf("Starting server on port: %d", addr)

	if err := http.ListenAndServe(addrStr, d.Router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
