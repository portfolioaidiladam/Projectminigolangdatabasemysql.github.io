// Package repository berisi implementasi akses data untuk entitas Comment
package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

// commentRepositoryImpl merupakan implementasi dari interface CommentRepository
// yang menangani operasi database untuk entitas Comment
type commentRepositoryImpl struct {
	DB *sql.DB
}

// NewCommentRepository membuat instance baru dari CommentRepository
// dengan menerima koneksi database sebagai parameter
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

// Insert menambahkan comment baru ke dalam database
// Menerima context untuk penanganan timeout dan comment yang akan disimpan
// Mengembalikan comment yang telah disimpan (dengan ID baru) dan error jika ada
func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

// FindById mencari comment berdasarkan ID yang diberikan
// Menerima context dan ID yang dicari
// Mengembalikan comment jika ditemukan dan error jika terjadi masalah atau data tidak ditemukan
func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// Scan data ke struct comment jika ditemukan
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// Kembalikan error jika data tidak ditemukan
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

// FindAll mengambil semua data comment dari database
// Menerima context untuk penanganan timeout
// Mengembalikan slice dari Comment dan error jika ada
func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
