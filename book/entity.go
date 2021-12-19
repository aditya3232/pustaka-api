package book

import "time"

// membuat tabel Book
// nanti namanya bakalan otomatis jadi books di database
// nanti nama column bakalan kecil semua, dan creted_at & updated_at
// jika mau tambah kolom didalam tabel books, tinggal tambahi aja disini. Ntar waktu di run akan otomatis nambah
type Book struct{
	ID int
	Title string
	Description string
	Price int
	Rating int
	Discount int
	CreatedAt time.Time
	UpdatedAt time.Time
}