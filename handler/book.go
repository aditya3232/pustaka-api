package handler // sesuaikan dengan nama folder
import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// menggunakan service didalam handler
type bookHandler struct{
	bookService book.Service
}
// menggunakan service didalam handler
func NewBookHandler(bookService book.Service) *bookHandler{
	return &bookHandler{bookService}
}

// // penulisan fungsi handler di folder ini agar dapat digunakan oleh main, maka nama fungsi diawali dengan huruf besar
// func (h *bookHandler) RootHandler(c *gin.Context) { 
// 	c.JSON(http.StatusOK, gin.H{
// 		"name": "Agung Setiawan",
// 		"bio":  "A Software Engineer & Content creator",
// 	})
// }

// // diberi (handler *bookHandler) agar menggunakan service
// func (h *bookHandler) HelloHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"title":    "Hello World",
// 		"subtitle": "Belajar Golang bareng Agung Setiawan",
// 	})
// }

// func (h *bookHandler) BooksHandler(c *gin.Context) {
// 	id := c.Param("id") // menangkap variabel id dari router menggunakan c.Param. Kemudian disimpan di variabel id
// 	title := c.Param("title")

// 	c.JSON(http.StatusOK, gin.H{"id": id, "title": title}) // kemudian ditampilkan di JSON
// }

// func (h *bookHandler) QueryHandler(c *gin.Context) {
// 	title := c.Query("title")
// 	price := c.Query("price")

// 	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
// }

// fungsi untuk handle custom response
// dibuat ini agar tidak menulis baris kode yang sama disetiap func
func convertBookResponse(b book.Book) book.BookResponse{
	return book.BookResponse{
			ID: b.ID,
			Title: b.Title,
			Price: b.Price,
			Description: b.Description,
			Rating: b.Rating,
			Discount: b.Discount,
		}
}


// Post Book
func (h *bookHandler) CreateBook(c *gin.Context){
	// variable bookRequest
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	// kondisi jika error
	if err != nil{
		//log.Fatal(err) // ini artinya jika error, server akan mati

		// menampilkan error di response dalam bentuk JSON ketika inputan tidak sesuai validasi. Dapat menerima lebih dari satu error
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return // didberikan return, agar jika ada error, tidak dilanjutkan kondisi yg dibawah
		//fmt.Println(err) // menampilkan informasi error apa, di terminal

	}

	// memanggil service
	book, err := h.bookService.Create(bookRequest)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return 
	}

	// kondisi jika tidak ada error akan menampilkan json
	c.JSON(http.StatusOK, gin.H{
		"data": convertBookResponse(book),
	})
}

// setiap fungsi di handler harus didaftarkan di routing pada main.go
// setiap fungsi di handler harus menggunakan awalan kapital, agar tidak dianggap private
func  (h *bookHandler) GetBooks(c *gin.Context){
	books, err := h.bookService.FindAll()
	// jika ada error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	// menerapkan BookResponse (custom response JSON)
	// hati2 disini ada perbedaan antara booksResponse, dgn bookResponse
	var booksResponse []book.BookResponse
	for _, b := range books {
		bookResponse := convertBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)

	}
	// kalau ada data bukunya, mengembalikan json
	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

// ambil buku by id
func (h *bookHandler) GetBook(c *gin.Context){
	// mengambil id dari routes
	idString := c.Param("id")
	// mmerubah string id, menjadi int
	id, _ := strconv.Atoi(idString)
	// panggil service FindByID
	b, err := h.bookService.FindByID(int(id))
	// jika ada error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	// menerapkan BookResponse (custom response JSON)
	bookResponse := convertBookResponse(b)
	// kalau ada data bukunya, mengembalikan json
	c.JSON(http.StatusOK, gin.H{
		// data yang diresponse adalah b => dari panggil service FindByID =>      b, err := h.bookService.FindByID(int(id))
		// namun karena menerapkan BookResponse, jadi yg data yg dikelaurkan adalah bookResponse 
		"data": bookResponse, 
	})
}

// Update Book
func (h *bookHandler) UpdateBook(c *gin.Context){
	// variable bookRequest
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	// kondisi jika error
	if err != nil{
		//log.Fatal(err) // ini artinya jika error, server akan mati

		// menampilkan error di response dalam bentuk JSON ketika inputan tidak sesuai validasi. Dapat menerima lebih dari satu error
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return // didberikan return, agar jika ada error, tidak dilanjutkan kondisi yg dibawah
		//fmt.Println(err) // menampilkan informasi error apa, di terminal

	}
	// mengambil id dari routes
	idString := c.Param("id")
	// mmerubah string id, menjadi int
	id, _ := strconv.Atoi(idString)
	// memanggil service Update (ada 2 parameter yaitu id, bookRequest)
	book, err := h.bookService.Update(id, bookRequest)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return 
	}

	// kondisi jika tidak ada error akan menampilkan json
	c.JSON(http.StatusOK, gin.H{
		"data": convertBookResponse(book),
	})
}

// hapus buku
func (h *bookHandler) DeleteBook(c *gin.Context){
	// mengambil id dari routes
	idString := c.Param("id")
	// mmerubah string id, menjadi int
	id, _ := strconv.Atoi(idString)
	// panggil service Delete
	b, err := h.bookService.Delete(int(id))
	// jika ada error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	// menerapkan BookResponse (custom response JSON)
	bookResponse := convertBookResponse(b)
	// kalau ada data bukunya, mengembalikan json
	c.JSON(http.StatusOK, gin.H{
		// data yang diresponse adalah b => dari panggil service FindByID =>      b, err := h.bookService.FindByID(int(id))
		// namun karena menerapkan BookResponse, jadi yg data yg dikelaurkan adalah bookResponse 
		"data": bookResponse, 
	})
}