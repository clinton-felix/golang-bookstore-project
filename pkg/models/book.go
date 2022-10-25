package models

import (
	"github.com/clinton-felix/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Creating a book struct
type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json: "author"`
	Publication string `json: "publication"`
}

// creating an init function that will help us initialize the database from the config file..
func init()  {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

/* creating the custom methods on the Book Struct */

// Create book method
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// Get All books method
// NB we are returning a slice bcos we want to return the list of books in the DB
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// Getting a Book by ID

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// Deleting a book by ID

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}