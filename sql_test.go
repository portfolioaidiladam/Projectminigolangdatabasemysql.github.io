// Package belajar_golang_database berisi implementasi pengujian untuk operasi database MySQL
package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// TestExecSql menguji fungsi untuk melakukan operasi INSERT ke database
// Best Practice: Selalu gunakan ExecContext daripada Exec biasa untuk mendukung timeout dan cancellation
func TestExecSql(t *testing.T) {
	// Dapatkan koneksi database
	db := GetConnection()
	// Best Practice: Selalu tutup koneksi database setelah selesai digunakan
	defer db.Close()

	// Best Practice: Gunakan context untuk mengontrol timeout dan cancellation
	ctx := context.Background()

	// Best Practice: Hindari hard-coded values dalam query SQL
	script := "INSERT INTO customer(id, name) VALUES('Aidil', 'Aidil')"
	_, err := db.ExecContext(ctx, script)
	// Best Practice: Selalu handle error yang mungkin terjadi
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

// TestQuerySql menguji fungsi untuk melakukan query SELECT sederhana
// Best Practice: Gunakan QueryContext untuk operasi SELECT
func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Best Practice: Tentukan kolom yang dibutuhkan, hindari SELECT *
	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	// Best Practice: Selalu tutup rows setelah selesai
	defer rows.Close()

	// Best Practice: Gunakan rows.Next() untuk iterasi hasil query
	for rows.Next() {
		var id, name string
		// Best Practice: Gunakan Scan untuk mengambil nilai dari rows
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

// TestQuerySqlComplex menguji query dengan berbagai tipe data
// Best Practice: Gunakan tipe data sql.NullXXX untuk menangani nilai NULL
func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}
}

// TestSqlInjection mendemonstrasikan kerentanan SQL Injection
// Best Practice: JANGAN PERNAH menggunakan string concatenation untuk query SQL
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

// TestSqlInjectionSafe mendemonstrasikan cara aman untuk menghindari SQL Injection
// Best Practice: Selalu gunakan parameter binding (?) untuk nilai dinamis
func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

// TestExecSqlParameter menguji penggunaan parameter dalam query INSERT
// Best Practice: Gunakan parameter binding untuk semua nilai yang diinput user
func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "aidil'; DROP TABLE user; #"
	password := "aidil"

	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

// TestAutoIncrement menguji penggunaan auto increment dan LastInsertId
// Best Practice: Gunakan LastInsertId untuk mendapatkan ID yang baru dibuat
func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "eko@gmail.com"
	comment := "Test komen"

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)
}

// TestPrepareStatement menguji penggunaan prepared statement
// Best Practice: Gunakan prepared statement untuk query yang dieksekusi berulang kali
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)
	}
}

// TestTransaction menguji penggunaan transaksi database
// Best Practice: Gunakan transaksi untuk operasi yang membutuhkan atomicity
func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err :=db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	// do transaction
	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}

