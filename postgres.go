package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() *Postgres {
	connStr := "user=moritz dbname=taskmanage password=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil
	}

	return &Postgres{
		db,
	}

}

func (psql *Postgres) Init() {
	query := `CREATE TABLE IF NOT EXISTS users(
		id serial PRIMARY KEY,
		username VARCHAR(50) UNIQUE,
		email TEXT UNIQUE,
		password TEXT,
		createdAt TIMESTAMP
	)`

	_, err := psql.db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

}

func (psql *Postgres) AddNewAccount(acc *Account) error {

	query := `INSERT INTO users (username, email, password, createdAt) VALUES($1, $2, $3, $4)`

	_, err := psql.db.Exec(query, acc.Username, acc.Email, acc.Password, acc.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (psql *Postgres) RemoveAccount(acc *Account) error {
	query := `DELETE FROM users WHERE username=$1 AND email=$2`
	_, err := psql.db.Exec(query, acc.Username, acc.Email)
	return err
}
func (psql *Postgres) GetAccountByName(name string) (*Account, error) {
	query := `SELECT * FROM users WHERE username=$1`
	rows, err := psql.db.Query(query, name)
	if err != nil {
		return nil, err
	}
	acc := &Account{}
	for rows.Next() {
		if err := rows.Scan(&acc.Username, &acc.Email, &acc.Password, &acc.CreatedAt); err != nil {
			return nil, err
		}

	}

	return acc, nil
}

func (psql *Postgres) GetAccountByEmail(email string) (*Account, error) {
	query := `SELECT * FROM users WHERE email=$1`
	rows, err := psql.db.Query(query, email)
	if err != nil {
		return nil, err
	}
	acc := &Account{}
	for rows.Next() {
		if err := rows.Scan(&acc.Username, &acc.Email, &acc.Password, &acc.CreatedAt); err != nil {
			return nil, err
		}

	}

	return acc, nil
}
