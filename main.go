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

	//route to add url
	http.HandleFunc("/add", handlers.AddURLHandler(urlService))
	//route to redirect
	http.HandleFunc("/", handlers.RedirectHandler(urlService))
	//route to get stats for all links
	http.HandleFunc("/stats", handlers.StatsHandler(urlService))
	//route to get stats by url
	http.HandleFunc("/stats/", handlers.ClicksHandler(urlService))

	log.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
