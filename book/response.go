package book

// response.go disini berfungsi untuk mengcustom key saat kita menjalankan route GetBooks, dimana sebelumnya mengikuti entity, sekarang mengikuti response.go
// selain mengcustom key, disini juga bisa mengcustom response json apa saja yang diharapkan keluar. dalam kasus ini, createdat & updatedat tidak dikeluarkan di response json
// setelah membuat BookResponse, terapkan didalam handler book.go
type BookResponse struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price int `json:"price"`
	Rating int `json:"rating"`
	Discount int `json:"discount"`
}