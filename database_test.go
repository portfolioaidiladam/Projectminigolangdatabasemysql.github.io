// Package belajar_golang_database berisi fungsi-fungsi testing untuk koneksi database MySQL
package belajar_golang_database

import (
	"database/sql" // package untuk menangani operasi SQL
	"testing"      // package untuk unit testing

	_ "github.com/go-sql-driver/mysql" // driver MySQL yang diimport secara blank untuk registrasi driver
)

// TestEmpty adalah fungsi test kosong sebagai placeholder
func TestEmpty(t *testing.T) {
	// Belum ada implementasi
}

// TestOpenConnection menguji koneksi ke database MySQL
func TestOpenConnection(t *testing.T) {
	// Membuka koneksi ke database MySQL
	// Format DSN: username:password@protocol(host:port)/dbname
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err) // Menghentikan program jika terjadi error saat koneksi
	}
	// Menutup koneksi database setelah fungsi selesai dijalankan
	defer db.Close()

	// Placeholder untuk penggunaan koneksi database
	// gunakan DB
}


