package entity

// Comment merepresentasikan struktur data untuk komentar dalam sistem
type Comment struct {
	// Id adalah identifier unik untuk setiap komentar
	Id int32

	// Email menyimpan alamat email dari pengguna yang membuat komentar
	Email string

	// Comment menyimpan isi pesan komentar dari pengguna
	Comment string
}
