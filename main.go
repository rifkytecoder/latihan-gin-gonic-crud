package main

import (
	"fmt"
	"lab-go-gin-crud/book"
	"lab-go-gin-crud/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// todo GORM deklarasi
	// <user:password> ... /nama_DB_api
	dsn := "root:admin@tcp(127.0.0.1:3306)/pustaka_api_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //testing !

	// Test connections
	if err != nil {
		log.Fatal("Database connetion ERROR")
	}
	fmt.Println("=============================================") //testing !
	fmt.Println(">>>>>>> Database connection SUCCEED <<<<<<<<<") //testing !
	fmt.Println("=============================================") //testing !

	//todo Migration
	db.AutoMigrate(&book.Book{}) // DDL create Gorm entity `book` di Mysql from pkg book

	// ? Layering
	bookRepository := book.NewRepository(db)            //Instance Layer Repository
	bookService := book.NewService(bookRepository)      //Instance Layer Service
	bookHandler := handlers.NewBookHandler(bookService) //Instance Layer Handler

	//todo Read all data Hardcode manual console data dari MySql
	// books, err := bookRepository.FindAll()
	// if err != nil {
	// 	fmt.Println("Error Finding book")
	// }
	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// }

	//todo Read by Id data Hardcode manual console data dari Mysql
	// book, err := bookRepository.FindByID(2)
	// if err != nil {
	// 	fmt.Println("Error Finding book")
	// }
	// fmt.Println("Title :", book.Title)

	// todo Create `layer repository` book Hardcode manual ke Mysql
	// book := book.Book{
	// 	Title:       "Gundam",
	// 	Description: "Perang Robot",
	// 	Price:       200000,
	// 	Rating:      5,
	// }
	// newBook, err := bookRepository.Create(book)
	// if err != nil {
	// 	fmt.Println("Error Creating Book")
	// }
	// fmt.Printf("New data book %v ", newBook)

	// todo Create `layer service` book Hardcode manual ke Mysql
	//proses ini seharusnya terjadi di handler/ postman `JSON`
	// bookRequest := book.BookRequest{
	// 	Title: "Crush Gear",
	// 	Price: "45000",
	// }
	//func Create yg di service parameternya BookRequest
	//bookService.Create(bookRequest)

	// ? CRUD
	// todo CREATE/INSERT data Hard code Manual ke Mysql
	// tiap kali server di running data bertambah/ter create !cuma testing
	// book := book.Book{}
	// book.Title = "Dragon Ball Z"
	// book.Price = 40000
	// book.Rating = 5
	// book.Description = "Mengumpulkan Bola Dragon ball"
	//
	// err = db.Create(&book).Error // DML insert `book` Gorm
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }
	//
	// todo READ data Hard code Manual ke Mysql
	// var book book.Book
	// //err = db.First(&book).Error // DML read `book` Gorm first data
	// err = db.First(&book, 2).Error // DML read `book` Gorm By ID
	// //err = db.Debug().Last(&book).Error // DML read `book` Gorm last data
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Finding book record")
	// 	fmt.Println("==========================")
	// }
	// fmt.Println("Title :", book.Title)
	// fmt.Printf("book object %v ", book)
	//
	// todo read all data Hardcode Manual ke Mysql
	// var books []book.Book
	// //err = db.Find(&books).Error // DML read `book` Gorm find All
	// err = db.Where("price = ?", 40000).Find(&books).Error // DML read `book` Gorm By same String/value
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Finding book record")
	// 	fmt.Println("==========================")
	// }
	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Printf("book object %v ", b)
	// }
	//
	//todo UPDATE data Hardcode Manual ke Mysql
	// var book book.Book
	// // pilih data yg mau di Update
	// err = db.Where("id = ?", 1).First(&book).Error // DML read `book` Gorm limit 1
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Finding book record")
	// 	fmt.Println("==========================")
	// }
	// // pilih Field apa yg mau di update dan ganti valuenya
	// book.Title = "One Piece (Revised Edition)"
	//
	// err = db.Save(&book).Error // Save perubahan ke database Mysql nya
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Updating book record")
	// 	fmt.Println("==========================")
	// }
	//
	//todo Delete data Hardcode Manual ke Mysql
	// var book book.Book
	// // pilih data yang mau dihapus
	// err = db.Where("id = ?", 3).First(&book).Error // DML read `book` Gorm
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Finding book record")
	// 	fmt.Println("==========================")
	// }
	// err = db.Delete(&book).Error // DML Delete data `book` di Mysql
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Deleting book record")
	// 	fmt.Println("==========================")
	// }
	//!========================================================================//

	// todo Gin Framework
	router := gin.Default()
	//path URL "Anonymous function"
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"messageA": "pong and pong",
		})
	})

	// todo Router GROUP VERSIONING
	v1 := router.Group("/v1")
	// Routing
	//!v1.GET("/hello", HelloHandler) //telah menggunakan groub
	//v1.GET("/hello", handlers.HelloHandler) //telah menggunakan pkg handlers

	//!router.GET("/books/:id/:tema", BooksHandler) before
	// router.GET("/books/:id/:tema", handlers.HelloHandler) //pkg handlers
	// router.GET("/query", handlers.QueryHandler)
	// router.POST("/books", handlers.PostBookHandler) // tanpa versioning

	// todo layer handlers
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)       //Layer Handlers
	v1.PUT("/books/:id", bookHandler.UpdateBook)    //Layer Handlers
	v1.DELETE("/books/:id", bookHandler.DeleteBook) //Layer Handlers

	router.Run(":8888") //default port 8080
}

// todo Handler "function"
// func HelloHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message":  "Welcome",
// 		"messageA": "Gin Gonic",
// 		"messageB": "Route Group path URL",
// 	})
// }

// todo Book Holder/contains information
// type BookResponse struct {
// 	Title string      `json:"title" binding:"required"`        //validate
// 	Price json.Number `json:"price" binding:"required,number"` //json string/int yg penting angka
// 	//SubTitle string `json:"sub_title"`
// }

// func PostBookHandler(c *gin.Context) {
// 	// membuat object bookResponse
// 	//var bookResponse BookResponse //before
// 	var bookResponse book.BookResponse //pkg book

// 	// if err := c.ShouldBindJSON(&bookResponse); err != nil { //cara satu
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

// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": errorMessages,
// 		})
// 		return

// 		//log.Fatal(err)
// 		// c.JSON(http.StatusBadRequest, err)
// 		// fmt.Println(err)
// 		// return
// 	}
// 	// jika success
// 	c.JSON(http.StatusOK, gin.H{
// 		"title": bookResponse.Title,
// 		"price": bookResponse.Price,
// 		//"sub_title": bookResponse.SubTitle,
// 	})
// }
