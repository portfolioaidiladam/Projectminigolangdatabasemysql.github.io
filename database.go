// Package belajar_golang_database menyediakan fungsi-fungsi untuk mengelola koneksi database MySQL
package belajar_golang_database

import (
	"database/sql"
	"time"
)

// GetConnection mengembalikan pointer ke sql.DB yang sudah dikonfigurasi
// Fungsi ini menginisialisasi koneksi ke database MySQL dengan pengaturan yang optimal
// Return: *sql.DB - objek koneksi database yang siap digunakan
func GetConnection() *sql.DB {
	// Membuat koneksi ke database MySQL
	// Format DSN: username:password@protocol(host:port)/dbname?param=value
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err) // Menghentikan program jika terjadi error saat koneksi
	}

	// SetMaxIdleConns menentukan jumlah maksimum koneksi idle yang disimpan di connection pool
	db.SetMaxIdleConns(10)
	
	// SetMaxOpenConns membatasi jumlah maksimum koneksi yang bisa dibuka secara bersamaan
	db.SetMaxOpenConns(100)
	
	// SetConnMaxIdleTime mengatur berapa lama koneksi bisa idle sebelum ditutup
	db.SetConnMaxIdleTime(5 * time.Minute)
	
	// SetConnMaxLifetime mengatur waktu maksimum sebuah koneksi bisa digunakan
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
