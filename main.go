package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    _ "my-gin-swagger-app/docs" // Replace 'ur package' with your module name
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
)

type Book struct {
    ID     string `json:"id"` //example used to show the value of ur request body on swagger
    Title  string `json:"title"`
    Author string `json:"author"`
}

type ErrorResponse struct {
    Message string `json:"message"`
}

type SuccessResponse struct {
    Message string `json:"message"`
}

type CreateBookRequest struct {
    Title  string `json:"title" example:"1984"`
    Author string `json:"author" example:"George Orwell"`
}

type UpdateBookRequest struct {
    Title  string `json:"title" example:"1984"`
    Author string `json:"author" example:"George Orwell"`
}

var books = []Book{
    {ID: "1", Title: "1984", Author: "George Orwell"},
    {ID: "2", Title: "Brave New World", Author: "Aldous Huxley"},
    {ID: "3", Title: "The Catcher in the Rye", Author: "J.D. Salinger"},
}

// @title Book API
// @version 1.0
// @description This is a sample server for managing books.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func main() {
    r := gin.Default()

    // Swagger setup
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.GET("/books", getBooks)
    r.GET("/books/:id", getBookByID)
    r.POST("/books", createBook)
    r.PUT("/books/:id", updateBook)
    r.DELETE("/books/:id", deleteBook)

    r.Run() // listen and serve on 0.0.0.0:8080
}

// getBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Produce json
// @Success 200 {array} Book
// @Router /books [get]
func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, gin.H{
		"message": "get all books successfully",
		"payload" : books,
	})
}

// getBookByID godoc
// @Summary Get a book by ID
// @Description Get details of a book by its ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} Book
// @Failure 404 {object} ErrorResponse
// @Router /books/{id} [get]
func getBookByID(c *gin.Context) {
    id := c.Param("id")
    for _, book := range books {
        if book.ID == id {
            c.IndentedJSON(http.StatusOK, book)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "Book not found",
	})
}

// createBook godoc
// @Summary Create a new book
// @Description Create a new book with the input payload
// @Tags books
// @Accept json
// @Produce json
// @Param book body CreateBookRequest true "Book to create"
// @Success 201 {object} Book
// @Router /books [post]
func createBook(c *gin.Context) {
    var newBookRequest CreateBookRequest

    if err := c.BindJSON(&newBookRequest); err != nil {
        return
    }

    // Auto-increment ID
    newID := strconv.Itoa(len(books) + 1)
    newBook := Book{
        ID:     newID,
        Title:  newBookRequest.Title,
        Author: newBookRequest.Author,
    }

    books = append(books, newBook)
    c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Create book successfully",
		"payload": newBook,
	})
}

// updateBook godoc
// @Summary Update a book
// @Description Update details of a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body UpdateBookRequest true "Updated book"
// @Success 200 {object} SuccessResponse
// @Failure 404 {object} ErrorResponse
// @Router /books/{id} [put]
func updateBook(c *gin.Context) {
    id := c.Param("id")
    var updatedBookRequest UpdateBookRequest

    if err := c.BindJSON(&updatedBookRequest); err != nil {
        return
    }

    for i, book := range books {
        if book.ID == id {
            books[i].Title = updatedBookRequest.Title
            books[i].Author = updatedBookRequest.Author
            c.IndentedJSON(http.StatusOK, SuccessResponse{Message: "book updated"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, ErrorResponse{Message: "book not found"})
}

// deleteBook godoc
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Param id path string true "Book ID"
// @Success 200 {object} SuccessResponse
// @Failure 404 {object} ErrorResponse
// @Router /books/{id} [delete]
func deleteBook(c *gin.Context) {
    id := c.Param("id")

    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            c.IndentedJSON(http.StatusOK, SuccessResponse{Message: "yrkkketed"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, ErrorResponse{Message: "book not found"})
}
