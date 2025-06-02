package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type User struct {
	Name      string
	BirthYear int64
}

func main() {
	db, err := newConn()

	if err != nil {
		fmt.Println("newDB.Error()", err)
		return
	}

	defer func() {
		_ = db.Close()
		fmt.Println("koneksi ditutup")
	}()

	// buat helper untuk insert user dengan pointer
	newName := func(a string) *string {
		return &a
	}
	birthYear := func(b int64) *int64 {
		return &b
	}

	if err := InsertUsers(db, []User{
		{
			Name:      *newName("Ahmad Suaib 2"),
			BirthYear: *birthYear(2003),
		},
		{
			Name:      *newName(""),
			BirthYear: *birthYear(2004),
		},
	}); err != nil {
		fmt.Println("InsertUsers.Error()", err)
		return
	}
}

func InsertUsers(db *sql.DB, users []User) error {
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("db.BeginTx %w", err)
	}

	// jika terjadi sesuai maka di rollback
	// jika tidak maka di commit
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
	}()

	for _, user := range users {
		_, err := tx.ExecContext(context.Background(), "INSERT INTO users(name, birth_year) VALUES($1, $2)", user.Name, user.BirthYear)
		if err != nil {
			return fmt.Errorf("tx.ExecContext %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("tx.Commit %w", err)
	}

	return nil
}

func newConn() (*pgx.Conn, error) {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5432",
		User:   url.UserPassword("root", "pass@word1"),
		Path:   "BelajarGoLang",
		RawQuery: url.Values{
			"sslmode": []string{"disable"},
		}.Encode(),
	}

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("sql.Open %w", err)
	}

	return db, nil
}
