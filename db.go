package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDatabase() {
	connectionString := "host=localhost user=postgres password=postgres dbname=rinha_backend port=5432 sslmode=disable"
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panicf("error: %s", err)
	}

	db = database
}

func CreateTable() {
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

func InsertTable(correlationId uuid.UUID, amount float64, status string) {
	command := `
		insert into payments(?, ?, ?, ?);
	`

	fmt.Println(command)
}
