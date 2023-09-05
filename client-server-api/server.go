package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

	db = persistence.NewDatabase(dbConn)

	http.HandleFunc("/cotacao", getUSDBRLExchangeRateHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getUSDBRLExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	c := httpclient.HttpClient{
		BaseURL: "https://economia.awesomeapi.com.br",
		Timeout: 200 * time.Millisecond,
	}

	var rateResp models.ExchangeRateResponse

	if err := c.Get("/json/last/USD-BRL", &rateResp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	stmt, err := db.Connection.Prepare("INSERT INTO rates (code, code_in, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(
		rateResp.USDBRL.Code,
		rateResp.USDBRL.CodeIn,
		rateResp.USDBRL.Name,
		rateResp.USDBRL.High,
		rateResp.USDBRL.Low,
		rateResp.USDBRL.VarBid,
		rateResp.USDBRL.PctChange,
		rateResp.USDBRL.Bid,
		rateResp.USDBRL.Ask,
		rateResp.USDBRL.Timestamp,
		rateResp.USDBRL.CreateDate,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rateResp.USDBRL)
}
