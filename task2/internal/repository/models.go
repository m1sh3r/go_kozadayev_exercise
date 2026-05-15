package repository

type Document struct {
	ID     int64
	Title  string
	UserID int64
}

type User struct {
	ID   int64
	Name string
}
