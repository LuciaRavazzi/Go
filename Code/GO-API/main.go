//------ Create an API in Go using the Gin framework.

// go mod init example/GO-API
// Obtain gin.
// go get github.com/gin-gonic/gin
// Run the API
// go run ./main.go

// The API stores a buch of books and extract some information.
package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// json is the chosen representation of data
// for the API.
// The capital letter defines a public field
// which allows external modules to see those fields.
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	// return the json version of the book.
	c.IndentedJSON(http.StatusOK, books)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id book."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not available."})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id book."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not available."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found.")
}

func main() {

	/* Main HTTP verbs:
	- POST: create.
	- GET: read
	- PUT: update
	- PATCH: modify.
	*/

	// router handles different routers.
	router := gin.Default()

	/*
	   Get request of our API.
	   To get data:
	   Power shell: curl http://localhost:8080/books
	   CMD: 		curl http://localhost:8080/books
	   curl -> Client URL.
	*/
	router.GET("/books", getBooks)

	// Return book whose id is given. The path must define the id.
	router.GET("/books/:id", bookById)

	/*
		   Post request.
		   It returns the new book which is added to the book collection.
		   CMD: curl http://localhost:8080/books -d @body.json --include
		  		-d: data. @body.json gathers data.
				--include: include additional information.
	*/
	router.POST("/books", createBook)

	// Check the id: curl localhost:8080/checkout?id=2 --request PATCH
	router.PATCH("/checkout", checkoutBook)
	// Add quantity.
	router.PATCH("/return", returnBook)

	router.Run("localhost:8080")

}
