package main

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

type Borrowing struct {
	BorrowingID int    `json:"borrowing_id"`
	MemberID    int    `json:"member_id"`
	BookID      int    `json:"book_id"`
	BorrowDate  string `json:"borrow_date"`
	ReturnDate  string `json:"return_date"`
	Status      string `json:"status"`
}

type Reservation struct {
	ReservationID   int    `json:"reservation_id"`
	MemberID        int    `json:"member_id"`
	BookID          int    `json:"book_id"`
	ReservationDate string `json:"reservation_date"`
	Status          string `json:"status"`
}

type BorrowRequest struct {
	MemberID   int    `json:"member_id"`
	BookID     int    `json:"book_id"`
	BorrowDate string `json:"borrow_date"`
}
