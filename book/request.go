package book // folder book mengurusi kode2 yg berhubungan dengan book. seperti struct,
import "encoding/json"

// FUNGSI DARI REQUEST.GO ADALAH VALIDASI UNTUK POST/UPDATE
// menangkap data dari inputan (post)
// samakan penulisannya dengan key json inputannya
// nulis variabel disini diawali dengan huruf besar
type BookRequest struct{
	Title string `json:"title" binding:"required"` // variable Title akan menangkap inputan dgn key title, dan variable Title harus diisi
	Price json.Number `json:"price" binding:"required,number"` // binding adalah validasi JSON. disini artinya Price harus diisi, dan harus berformat number
	// angka boleh pakai int, boleh json.Number, kalau pakai json.Number numbernya boleh dimasukkan kedalam string atau tidak didalam string, tetap akan dikonvert ke number
	// semisal tidak bisa pakai json.Number, pakai int sebenernya juga gak masalah
	Description string `json:"description" binding:"required"`
	Rating json.Number `json:"rating" binding:"required,number"`
	Discount json.Number `json:"discount" binding:"required,number"`
}