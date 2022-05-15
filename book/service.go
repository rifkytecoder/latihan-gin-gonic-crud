package book

// todo Layer Service bertanggung jawab berhubungan Bisnis logic
type IService interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

//! Repository <- Service
type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FindAll() //todo cara mudah
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()
	//mapping ke entity Book
	book := Book{
		// parameter BookRequest Field
		Title:       bookRequest.Title,
		Price:       int(price), //convert to 32 krna di Book fieldnya int
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}
	//krna func create dri repo mngunkn Book entity `jdi ter Create di entity Book`
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	// parameter BookRequest Field
	book.Title = bookRequest.Title
	book.Price = int(price) //convert to 32 krna di Book fieldnya int
	book.Description = bookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	//krna func create dri repo mngunkn Book entity `jdi ter Create di entity Book`
	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	//krna func create dri repo mngunkn Book entity `jdi ter Create di entity Book`
	newBook, err := s.repository.Delete(book)
	return newBook, err
}
