package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/ngavinsir/api-supply-demand-covid19/handler"
)

//go:generate sqlboiler --wipe psql

func main() {
	router := chi.NewRouter()

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	log.Println("connected to db")

	api := handler.NewAPI(db)
	router.Mount("/api/v1", api.Router())

	port := ":4040"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = ":" + envPort
	}

	log.Printf("Server started on %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}