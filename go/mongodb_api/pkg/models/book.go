package models

import (
	"github.com/kamva/mgm/v3"
	"github.com/vyrekxd/mongodb_api/pkg/config"
	"gopkg.in/mgo.v2/bson"
)

var collection *mgm.Collection

type Book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name"`
	Author           string `json:"author"`
	Publication      string `json:"publication"`
}

func init() {
	config.Connect()

	collection = mgm.Coll(&Book{})
}

func (b *Book) CreateBook() *Book {
	collection.Create(b)

	return b
}

func GetAllBooks() []Book {
	Books := []Book{}

	err := collection.SimpleFind(&Books, bson.M{})

	if err != nil {
		panic(err)
	}

	return Books
}

func GetBookById(Id string) (*Book, *mgm.Collection) {
	var getBook Book

	collection.FindByID(Id, &getBook)

	return &getBook, collection
}

func DeleteBook(Id string) *Book {
	book, _ := GetBookById(Id)

	collection.Delete(book)

	return book
}
