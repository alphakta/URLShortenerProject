package main

import (
	"log"
	"net/http"

	"github.com/alphakta/URLShortenerProject/database"
	"github.com/alphakta/URLShortenerProject/domain/url"
	"github.com/alphakta/URLShortenerProject/handlers"
	urlRepository "github.com/alphakta/URLShortenerProject/repository/url"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.InitDB("root:@tcp(localhost:3306)/shorturl")
	defer db.Close()

	address := "http://localhost:8080/"
	repo := urlRepository.NewMySQLRepository(db)
	urlService := url.NewService(repo)

	http.HandleFunc("/add", handlers.AddURLHandler(urlService))
	http.HandleFunc("/", handlers.RedirectHandler(urlService))
	http.HandleFunc("/stats", handlers.StatsHandler(urlService))
	http.HandleFunc("/stats/", handlers.ClicksHandler(urlService))

	log.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
