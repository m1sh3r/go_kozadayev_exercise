package repository

type Car struct {
	ID      int64
	Brand   string
	OwnerID int64
}

type Owner struct {
	ID   int64
	Name string
}

type ServiceRecord struct {
	ID          int64
	CarID       int64
	Description string
	Date        int64
}
