package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mathcale/goexpert-course/client-server-api/internal/models"
	"github.com/mathcale/goexpert-course/client-server-api/internal/pkg/httpclient"
	"github.com/mathcale/goexpert-course/client-server-api/internal/pkg/persistence"
)

func main() {
	log.SetPrefix("[CLIENT] ")

	c := httpclient.NewHttpClient("http://localhost:8080", 300*time.Millisecond)

	var r models.ExchangeRate

	log.Println("Fetching exchange rate from server...")

	if err := c.Get("/cotacao", &r); err != nil {
		log.Fatalf("Error fetching rate: %s", err.Error())
	}

	log.Println("Saving rate on file...")

	f := persistence.NewFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY)

	if err := f.Write([]byte(fmt.Sprintf("Dólar: %s\n", r.Bid))); err != nil {
		log.Fatalf("Error saving rate: %s", err.Error())
	}
}
