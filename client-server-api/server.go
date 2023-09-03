package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mathcale/goexpert-course/client-server-api/models"
	"github.com/mathcale/goexpert-course/client-server-api/pkg/httpclient"
)

func main() {
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rateResp.USDBRL)
}
