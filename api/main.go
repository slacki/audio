package main

import (
	"api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/rs/cors"
)

func main() {
	db := initDb()
	// migrateDb(db)

	env := &handlers.Env{
		UploadPath: "./_uploads",
		DB:         db,
	}

	r := mux.NewRouter()
	apiV1 := r.PathPrefix("/v1").Subrouter()
	apiV1.Handle("/upload", handlers.Handler{Env: env, H: handlers.HandleUpload}).Methods("POST")
	apiV1.Handle("/info/{id:[a-zA-Z0-9]+}", handlers.Handler{Env: env, H: handlers.HandleInfo}).Methods("GET")
	apiV1.Handle("/contact", handlers.Handler{Env: env, H: handlers.HandleContact}).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://audio.slacki.io", "http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	})
	handler := c.Handler(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    "0.0.0.0:8081",
	}

	log.Fatal(srv.ListenAndServe())
}

func initDb() *scribble.Driver {
	db, err := scribble.New("./_data", nil)
	if err != nil {
		log.Fatalln("Database error:", err)
	}

	return db
}
