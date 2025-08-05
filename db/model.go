package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TestData struct {
	ID         int       `db:"id"`
	CustomerId string    `db:"customer_id"`
	ProductId  string    `db:"product_id"`
	Amount     int       `db:"amount"`
	Date       time.Time `db:"date"`
}

type DB struct {
	Conn *sqlx.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{Conn: db}, nil
}

func (d *DB) AddPurchase(p TestData) error {
	_, err := d.Conn.Exec(`
		INSERT INTO testdata (customer_id, product_id, amount)
		VALUES ($1, $2, $3)
	`, p.CustomerId, p.ProductId, p.Amount)
	return err
}

func (d *DB) GetPurchases(customerId string) ([]TestData, error) {
	var testdata []TestData
	err := d.Conn.Select(&testdata, `
		SELECT * FROM testdata WHERE customer_id=$1 ORDER BY date DESC 
	`, customerId)
	return testdata, err
}
