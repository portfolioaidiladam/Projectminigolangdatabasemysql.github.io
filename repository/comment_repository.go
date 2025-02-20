package repository

import (
	"belajar-golang-database/entity"
	"context"
)

// CommentRepository adalah interface yang mendefinisikan kontrak untuk operasi CRUD pada entitas Comment
type CommentRepository interface {
	// Insert menyimpan data comment baru ke dalam database
	// Parameter:
	//   - ctx: context untuk manajemen request
	//   - comment: entity.Comment yang akan disimpan
	// Returns:
	//   - entity.Comment yang telah disimpan
	//   - error jika terjadi kesalahan
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)

	// FindById mencari comment berdasarkan ID
	// Parameter:
	//   - ctx: context untuk manajemen request
	//   - id: primary key dari comment yang dicari
	// Returns:
	//   - entity.Comment yang ditemukan
	//   - error jika terjadi kesalahan atau data tidak ditemukan
	FindById(ctx context.Context, id int32) (entity.Comment, error)

	// FindAll mengambil semua data comment dari database
	// Parameter:
	//   - ctx: context untuk manajemen request
	// Returns:
	//   - slice dari entity.Comment yang berisi semua data comment
	//   - error jika terjadi kesalahan
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
