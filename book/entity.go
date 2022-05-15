package book

import "time"

// ? Deklarasi model untuk ke DB tablenya `Mysql` `dari dalam ke luar from code to mysql`
// Create Book table
type Book struct {
	// Nama kolum/field
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	Discount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
