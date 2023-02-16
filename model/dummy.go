package model

import (
	"database/sql"
	"log"
	"time"
)

// Dummy example
type Dummy struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"desktop chair"`
	Price     float64   `json:"price" example:"299.99"`
	CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" swaggerignore:"true"`
}

func (d *Dummy) QCreateDummy(db *sql.DB) error {
	d.CreatedAt = time.Now()
	d.UpdatedAt = d.CreatedAt
	err := db.QueryRow(
		"INSERT INTO dummy(name, price, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING id",
		d.Name, d.Price, d.CreatedAt, d.UpdatedAt).Scan(&d.ID)

	if err != nil {
		return err
	}

	return nil
}

func (d *Dummy) QGetDummies(db *sql.DB, start int, count int) ([]Dummy, error) {
	var dummies []Dummy
	rows, err := db.Query(
		"SELECT * FROM dummy LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	for rows.Next() {
		var dummy Dummy
		if err = rows.Scan(&dummy.ID, &dummy.Name, &dummy.Price,
			&dummy.CreatedAt, &dummy.UpdatedAt); err != nil {
			return nil, err
		}
		dummies = append(dummies, dummy)
	}

	return dummies, nil
}

func (d *Dummy) QGetDummy(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM dummy WHERE id=$1",
		d.ID).Scan(&d.ID, &d.Name, &d.Price, &d.CreatedAt, &d.UpdatedAt)
}
