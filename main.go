package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amarquezmazzeo/bootdev-go-server/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	// db setup
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("could not open db: %s", err)
	}
	dbQueries := database.New(db)

	mux := http.NewServeMux()
	apiCfg := &apiConfig{dbQueries: dbQueries}
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /api/healthz", handlerHealth)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /api/validate_chirp", handlerValidate)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("starting serv")
	log.Fatal(server.ListenAndServe())

}
