package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

// setelah install Go
// jgn lupa jalankan diterminal => go get -u github.com/gin-gonic/gin
// ganti user:pass, dbname
func main() {
	// =================================================================================================
	// KONEKSI DATABASE
	// =================================================================================================
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	  // informasi jika ada error
	  if err != nil{
		  log.Fatal("db connection error")
		}
	// informasi jika sukses konek database
	fmt.Println("Database connection succeeded")
		
	// migration tabel Book
	db.AutoMigrate(&book.Book{})
	// =================================================================================================
	// REPOSITORY 
	// =================================================================================================
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	// find all
	// books, err := bookRepository.FindAll()
	// for _, book := range books{
	// 	fmt.Println("Title :", book.Title)
	// }

	// find by id
	// book, err := bookRepository.FindByID(2)
	// fmt.Println("Title :", book.Title)

	// create
	// book := book.Book{
	// 	Title: "$100 StartUp",
	// 	Description: "Good book",
	// 	Price: 95000,
	// 	Rating: 4,
	// 	Discount: 0,
	// }
	// bookRepository.Create(book)

	// create pakai service.
	// bookRequest := book.BookRequest{
	// 	Title: "Gundam",
	// 	Price: "200000",
	// }
	// bookService.Create(bookRequest)

	// =================================================================================================
	// ROUTE
	// =================================================================================================
	// definisi fungsi router
	router := gin.Default()

	// router group. Biasanya digunakan untuk mengelompokkan route
	// nanti disetiap route standar router.GET diganti v1.GET
	v1 := router.Group("/v1")
	
	// mengirimkan json ketika akses localhost:8080 // GET ("/")
	// jgn lupa jalankan di terminal go run main.go
	// v1.GET("/", bookHandler.RootHandler)
 
	// setiap nambahi router, jgn lupa jalankan ulang server
	// v1.GET("/hello", bookHandler.HelloHandler)

	// route dengan parameter (Path Variable)
	// v1.GET("/books/:id/:title", bookHandler.BooksHandler)

	// route dengan key (Query String)
	// localhost:8080/query?title=bumi manusia
	// localhost:8080/query?price=40&title=cantik itu luka
	// v1.GET("/query", bookHandler.QueryHandler)

	// route post
	v1.POST("/books", bookHandler.CreateBook)
	// route getBooks
	// ketika route getBooks dijalankan di postman, yg dikembalikan keynya menyesuaikan dengan entity.go (seperti model di laravel)
	// bisa diubah keynya dengan cara membuat struct yg khusus mewakili data json response. buat file nya di folder book (package book) dengan nama response.go
	v1.GET("/books", bookHandler.GetBooks)
	// route FindByID
	// localhost:8888/v1/book/2
	v1.GET("/book/:id", bookHandler.GetBook)
	// route update
	v1.PUT("/book/:id", bookHandler.UpdateBook)
	// route delete
	v1.DELETE("/book/:id", bookHandler.DeleteBook)




	// router.Run(":8888") // custom port masukkan kedalam ()
	router.Run("localhost:8888") // biar tidak memunculkan firewall
}


