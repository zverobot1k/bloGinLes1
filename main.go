package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID    int    `json:"book_id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
	Category  string `json:"category"`
	Available bool   `json:"available"`
}

type Member struct {
	MemberID  int    `json:"member_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	JoinDate  string `json:"join_date"`
}

type Reservation struct {
	ReservationID   int    `json:"reservation_id"`
	MemberID        int    `json:"member_id"`
	BookID          int    `json:"book_id"`
	ReservationDate string `json:"reservation_date"`
	Status          string `json:"status"`
}

var books []Book
var members []Member
var reservations []Reservation

func main() {
	books = append(books, Book{BookID: 0, Title: "Go Programming", Author: "Maksim Blokhin", Publisher: "Tech", Year: 2022, Category: "Programming", Available: true})
	members = append(members, Member{MemberID: 0, FirstName: "Maksim", LastName: "Blokhin", Phone: "1234567890", Email: "maksim@example.com", Address: "123st Mira", JoinDate: "2022-01-01"})
	reservations = append(reservations, Reservation{ReservationID: 0, MemberID: 0, BookID: 0, ReservationDate: "11.08.2024", Status: "Active"})

	r := gin.Default()

	r.GET("/books", getBooks)
	r.GET("/books/:id", getBookByID)
	r.POST("/books", createBook)
	r.PUT("/books/:id", updateBook) //!!!
	r.DELETE("/books/:id", deleteBook)

	r.GET("/members", getMembers)
	r.GET("/members/:id", getMemberByID)
	r.POST("/members", createMember)
	r.PUT("/members/:id", updateMember) //!!!
	r.DELETE("/members/:id", deleteMember)

	r.GET("/reservations", getReservations)
	r.GET("/reservations/:id", getReservationByID)
	r.POST("/reservations", createReservation)
	r.PUT("/reservations/:id", updateReservation) //!!!
	r.DELETE("/reservations/:id", deleteReservation)

	r.Run(":8080")
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, book := range books {
		if book.BookID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func createBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, book := range books {
		if book.BookID == id {
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func deleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, book := range books {
		if book.BookID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func getMembers(c *gin.Context) {
	c.JSON(http.StatusOK, members)
}

func getMemberByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, member := range members {
		if member.MemberID == id {
			c.JSON(http.StatusOK, member)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Member not found"})
}

func createMember(c *gin.Context) {
	var newMember Member
	if err := c.ShouldBindJSON(&newMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	members = append(members, newMember)
	c.JSON(http.StatusCreated, newMember)
}

func updateMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedMember Member
	if err := c.ShouldBindJSON(&updatedMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, member := range members {
		if member.MemberID == id {
			members[i] = updatedMember
			c.JSON(http.StatusOK, updatedMember)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Member not found"})
}

func deleteMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, member := range members {
		if member.MemberID == id {
			members = append(members[:i], members[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Member deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Member not found"})
}

func getReservations(c *gin.Context) {
	c.JSON(http.StatusOK, reservations)
}

func getReservationByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, reservation := range reservations {
		if reservation.ReservationID == id {
			c.JSON(http.StatusOK, reservation)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Reservation not found"})
}

func createReservation(c *gin.Context) {
	var newReservation Reservation
	if err := c.ShouldBindJSON(&newReservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reservations = append(reservations, newReservation)
	c.JSON(http.StatusCreated, newReservation)
}

func updateReservation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedReservation Reservation
	if err := c.ShouldBindJSON(&updatedReservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, reservation := range reservations {
		if reservation.ReservationID == id {
			reservations[i] = updatedReservation
			c.JSON(http.StatusOK, updatedReservation)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Reservation not found"})
}

func deleteReservation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, reservation := range reservations {
		if reservation.ReservationID == id {
			reservations = append(reservations[:i], reservations[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Reservation deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Reservation not found"})
}
