package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
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
		fmt.Println("sql.Open()", err)
		return
	}

	// untuk menutup koneksi
	defer func() {
		_ = db.Close()
		fmt.Println("Koneksi ditutup")
	}()

	// untuk cek koneksi
	err = db.PingContext(context.Background())
	if err != nil {
		fmt.Println("db.PingContext()", err)
		return
	}
	// ambil data dari database
	row := db.QueryRowContext(context.Background(), "SELECT birth_year FROM users")

	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext()", err)
		return
	}

	var birth_year int64
	err = row.Scan(&birth_year)
	if err != nil {
		fmt.Println("row.Scan()", err)
		return
	}

	fmt.Println("birth_year:", birth_year)

	// insert data ke database
	// _, err = db.ExecContext(context.Background(), "INSERT INTO users (name, birth_year) VALUES ('joko', '1901')")
	// if err != nil {
	//	fmt.Println("db.ExecContext()", err)
	// }

	// insert data ke database dengan prepared statement
	//name := "joko tingkir"
	//birth_year = 1902

	//_, err = db.ExecContext(context.Background(), "INSERT INTO users (name, birth_year) VALUES ($1, $2)", name, birth_year)
	//if err != nil {
	//	fmt.Println("db.ExecContext()", err)
	//}

	// Coba QueryContext
	rows, err := db.QueryContext(context.Background(), "SELECT name, birth_year FROM users")
	if err != nil {
		fmt.Println("db.QueryContext()", err)
		return
	}

	// karena rows bertindak seperti iterator, maka kita harus menutup rows
	defer func() {
		_ = rows.Close()
		fmt.Println("rows ditutup")
	}()

	if rows.Err() != nil {
		fmt.Println("row.Err()", err)
		return
	}

	// extract data dari rows, seperti DataReader di .Net
	for rows.Next() {
		var name string
		var birth_year int64

		if err := rows.Scan(&name, &birth_year); err != nil {
			fmt.Println("rows.Scan()", err)
			return
		}

		fmt.Println("name:", name, "\tbirth_year:", birth_year)
	}
}
