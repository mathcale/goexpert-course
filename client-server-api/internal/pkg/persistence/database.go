package persistence

import (
	"context"
	"database/sql"
	"time"

	"github.com/mathcale/goexpert-course/client-server-api/internal/models"
)

type Database struct {
	connection *sql.DB
}

func NewDatabase(connection *sql.DB) *Database {
	return &Database{
		connection: connection,
	}
}

func (d *Database) CreateRate(rate models.ExchangeRate) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	stmt, err := d.connection.Prepare("INSERT INTO rates (code, code_in, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(
		ctx,
		rate.Code,
		rate.CodeIn,
		rate.Name,
		rate.High,
		rate.Low,
		rate.VarBid,
		rate.PctChange,
		rate.Bid,
		rate.Ask,
		rate.Timestamp,
		rate.CreateDate,
	)

	if err != nil {
		return err
	}

	return nil
}
