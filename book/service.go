package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error) // disini create pakai BookRequest, karena data yg dikirim adalah json
	Update(ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository // yg dipakai Repository interface (dari repository.go)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err

	// return s.repository.FindAll() //ini cara cepat tanpa return error. Sama aja seperti yg diatas sebenarnya
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err

}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	// mengubah price jadi int
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	// simpan di struct Book, di entity.go
	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	// mencari by ID dulu
	// disini tidak perlu mengembalikan eror
	book, _ := s.repository.FindByID(ID)
	// mengubah price jadi int
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	// simpan di struct Book, di entity.go
	// atau maksudnya simpan di database
	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	// mencari by ID dulu
	// disini tidak perlu mengembalikan eror
	book, _ := s.repository.FindByID(ID)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}