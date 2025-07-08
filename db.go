package main

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() {
	connectionString := "host=localhost user=postgres password=postgres dbname=rinha_backend port=5432 sslmode=disable"
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panicf("error: %s", err)
	}

	db = database
	createTable()
}

func createTable() {
	command := `
	create table if not exists payments(
		correlation_id uuid 			   primary key,
		amount 		   numeric(14, 2) not null,
		processor 	   varchar(10)    not null, -- 'default' / 'fallback'
		status    	   varchar(10)    not null, -- 'processed' / 'failed'
		processed_at   TIMESTAMPTZ    not null  default now()
	);`

	_, err := db.Exec(command)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertTable(correlationId uuid.UUID, amount float64, processor, status string) {
	command := `
		insert into payments (correlation_id, amount, processor, status)
		values ($1, $2, $3, $4)
	`

	_, err := db.Exec(command, correlationId, amount, processor, status)
	if err != nil {
		log.Printf("error: %s\n", err)
	}
}
