package handlers

import (
	"fmt"
	"lab-go-gin-crud/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// todo Handler "function"
// func HelloHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message":  "Welcome",
// 		"messageA": "Gin Gonic",
// 		"messageB": "Route Group path URL",
// 		"messageC": "Route Group path URL handlers",
// 	})
// }
// string ID `localhost:8888/books/20/anime`
// func BooksHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	tema := c.Param("tema")
// 	c.JSON(http.StatusOK, gin.H{
// 		"id":   id,
// 		"tema": tema,
// 	})
// }
// string Query `localhost:8888/query?title=hotgame&price=50000`
// func QueryHandler(c *gin.Context) {
// 	title := c.Query("title")
// 	price := c.Query("price")
//
// 	c.JSON(http.StatusOK, gin.H{
// 		"title": title,
// 		"price": price,
// 	})
// }

type bookHandler struct {
	bookService book.IService
}

func NewBookHandler(bookService book.IService) *bookHandler {
	return &bookHandler{bookService}
}

// todo Post handler `tanpa layering`
// func PostBookHandler(c *gin.Context) {
// 	// membuat object bookRequest
// 	//var bookResponse BookRequest //before
// 	var bookResponse book.BookRequest //pkg book
//
// 	// if err := c.ShouldBindJSON(&bookRequest); err != nil { //cara satu
// 	err := c.ShouldBindJSON(&bookResponse)
// 	// Membuat Try and Catch
// 	// Jika error
// 	if err != nil {
// 		// Pesan Error
// 		errorMessages := []string{}
// 		for _, e := range err.(validator.ValidationErrors) {
// 			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
// 			errorMessages = append(errorMessages, errorMessage)
// 			//c.JSON(http.StatusBadRequest, errorMessage) //tanpa slice
// 			//return
// 		}
//
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": errorMessages,
// 		})
// 		return
//
// 		//log.Fatal(err)
// 		// c.JSON(http.StatusBadRequest, err)
// 		// fmt.Println(err)
// 		// return
// 	}
// 	// jika success
// 	c.JSON(http.StatusOK, gin.H{
// 		"title": bookRequest.Title,
// 		"price": bookRequest.Price,
// 		//"sub_title": bookResponse.SubTitle,
// 	})
// }

// todo handler findAll book
func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	// todo Book entity diubah ke BookResponse untuk hasil JSON nya
	var booksResponse []book.BookResponse
	for _, b := range books {
		// bookResponse := book.BookResponse{
		// 	ID:          b.ID,
		// 	Title:       b.Title,
		// 	Price:       b.Price,
		// 	Description: b.Description,
		// 	Rating:      b.Rating,
		// 	Discount:    b.Discount,
		// }
		bookResponse := convertToBookResponse(b)
		// memasukkan data saat di looping
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		// "data": books, //saat masih menggunakan Book Entity
		"data": booksResponse, // menggunakan BookResponse
	})
}

//todo handler findByID book
func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	//book, err := h.bookService.FindByID(int(id)) //tanpa BookResponse
	b, err := h.bookService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	// bookResponse := book.BookResponse{
	// 	ID:          b.ID,
	// 	Title:       b.Title,
	// 	Price:       b.Price,
	// 	Description: b.Description,
	// 	Rating:      b.Rating,
	// 	Discount:    b.Discount,
	// }
	bookResponse := convertToBookResponse(b)
	c.JSON(http.StatusOK, gin.H{
		// "data": book, //saat masih menggunakan Book Entity
		"data": bookResponse,
	})
}

// todo Post/Create dgn Layer Handlers/ Created
func (h *bookHandler) CreateBook(c *gin.Context) {
	// membuat object bookRequest
	//var bookResponse BookRequest //before
	var bookRequest book.BookRequest //pkg book

	// if err := c.ShouldBindJSON(&bookRequest); err != nil { //cara satu
	err := c.ShouldBindJSON(&bookRequest)
	// Membuat Try and Catch
	// Jika error
	if err != nil {
		// Pesan Error
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			//c.JSON(http.StatusBadRequest, errorMessage) //tanpa slice
			//return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

		//log.Fatal(err)
		// c.JSON(http.StatusBadRequest, err)
		// fmt.Println(err)
		// return
	}

	//layer handlers create
	book, err := h.bookService.Create(bookRequest) // create data dgn format BookRequest Field
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	// jika success
	c.JSON(http.StatusOK, gin.H{
		// "title": bookRequest.Title,
		// "price": bookRequest.Price,
		////"sub_title": bookResponse.SubTitle,
		//"data": book, // tanpa convertToBookResponse
		"data": convertToBookResponse(book),
	})
}

// todo handler Update Layer Handler
func (h *bookHandler) UpdateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Update(id, bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{

		// "data": book,
		"data": convertToBookResponse(book),
	})
}

// todo Delete Handler layer
func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	//book, err := h.bookService.FindByID(int(id)) //tanpa BookResponse
	b, err := h.bookService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	// bookResponse := book.BookResponse{
	// 	ID:          b.ID,
	// 	Title:       b.Title,
	// 	Price:       b.Price,
	// 	Description: b.Description,
	// 	Rating:      b.Rating,
	// 	Discount:    b.Discount,
	// }
	bookResponse := convertToBookResponse(b)
	c.JSON(http.StatusOK, gin.H{
		// "data": book, //saat masih menggunakan Book Entity
		"data": bookResponse,
	})
}

// todo fungsi Converter Response
func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}
