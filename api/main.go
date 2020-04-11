package main

import (
	"api/handlers"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// db := initDb()
	// migrateDb(db)

	env := &handlers.Env{
		UploadPath: "./uploads",
	}

	r := mux.NewRouter()
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	// regular routes
	apiV1.Handle("/upload", handlers.Handler{Env: env, H: handlers.HandleUpload}).Methods("POST")

	srv := &http.Server{
		Handler:      cors.Default().Handler(r),
		Addr:         "0.0.0.0:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// func initDb() *sqlx.DB {
// 	dns := os.Getenv("DATABASE_CONN")
// 	if len(dns) == 0 {
// 		panic("Database not configured, please set DATABASE_CONN environment variable.")
// 	}
// 	db, err := sqlx.Open("mysql", dns)
// 	if err != nil {
// 		log.Fatalf("could not connect to the MySQL database... %v", err)
// 	}

// 	if err := db.Ping(); err != nil {
// 		log.Fatalf("could not ping DB... %v", err)
// 	}

// 	return db
// }

// func migrateDb(db *sqlx.DB) {
// 	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
// 	if err != nil {
// 		log.Fatalf("could not start sql migration... %v", err)
// 	}

// 	m, err := migrate.NewWithDatabaseInstance(
// 		fmt.Sprintf("file://%s", "migrations"), // file://path/to/directory
// 		"mysql", driver)

// 	if err != nil {
// 		log.Fatalf("migration failed... %v", err)
// 	}

// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatalf("An error occurred while syncing the database.. %v", err)
// 	}

// 	log.Println("Database migrated")
// }
