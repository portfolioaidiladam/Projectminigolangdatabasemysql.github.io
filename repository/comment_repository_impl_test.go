// Package repository berisi implementasi pengujian untuk repository komentar
package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// TestCommentInsert menguji fungsi Insert pada CommentRepository
// Test ini memastikan bahwa komentar baru dapat disimpan ke database dengan benar
func TestCommentInsert(t *testing.T) {
	// Inisialisasi repository dengan koneksi database
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	// Membuat context untuk operasi database
	ctx := context.Background()
	
	// Menyiapkan data komentar untuk pengujian
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	// Mencoba melakukan insert komentar
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		t.Fatalf("gagal melakukan insert komentar: %v", err)
	}

	// Menampilkan hasil insert
	fmt.Println(result)
}

// TestFindById menguji fungsi FindById pada CommentRepository
// Test ini memastikan bahwa komentar dapat ditemukan berdasarkan ID
func TestFindById(t *testing.T) {
	// Inisialisasi repository dengan koneksi database
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	// Mencoba mencari komentar berdasarkan ID
	comment, err := commentRepository.FindById(context.Background(), 24)
	if err != nil {
		t.Fatalf("gagal mencari komentar berdasarkan ID: %v", err)
	}

	// Menampilkan komentar yang ditemukan
	fmt.Println(comment)
}

// TestFindAll menguji fungsi FindAll pada CommentRepository
// Test ini memastikan bahwa semua komentar dapat diambil dari database
func TestFindAll(t *testing.T) {
	// Inisialisasi repository dengan koneksi database
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	// Mencoba mengambil semua komentar
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		t.Fatalf("gagal mengambil semua komentar: %v", err)
	}

	// Menampilkan semua komentar yang ditemukan
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
