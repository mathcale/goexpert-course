package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mathcale/goexpert-course/client-server-api/models"
	"github.com/mathcale/goexpert-course/client-server-api/pkg/httpclient"
	"github.com/mathcale/goexpert-course/client-server-api/pkg/persistence"
)

var db *persistence.Database

func main() {
	dbConn, err := sql.Open("sqlite3", "./rates.db")

	if err != nil {
		panic(err)
	}

	log.SetPrefix("[SERVER] ")

	db = persistence.NewDatabase(dbConn)

	http.HandleFunc("/cotacao", getUSDBRLExchangeRateHandler)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getUSDBRLExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	c := httpclient.HttpClient{
		BaseURL: "https://economia.awesomeapi.com.br",
		Timeout: 200 * time.Millisecond,
	}

	var rateResp models.ExchangeRateResponse

	log.Println("Fetching exchange rate from API...")

	if err := c.Get("/json/last/USD-BRL", &rateResp); err != nil {
		log.Printf("Error fetching rate: %s", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	log.Println("Saving rate on database...")

	if err := db.CreateRate(rateResp.USDBRL); err != nil {
		log.Printf("Error saving rate: %s", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rateResp.USDBRL)
}
