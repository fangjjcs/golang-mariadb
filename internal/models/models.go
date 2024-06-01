package models

import (
	"database/sql"
)

// Models is the wrapper for all models
type Models struct {
	DB DBModel
}

// NewModels returns a model type with database connection pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// User is the type for users
type User struct {
	ID         int
	Name       string
	DeptId     string
	EmployeeID string
}

// Create Menu Request
type CreateMenu struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Memo       string `json:"memo"`
	FileString string `json:"fileString"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

// Open Menu Request
type OpenMenu struct {
	ID         int    `json:"id"`
	CreateUser string `json:"createUser"`
	CreateDept string `json:"createDept"`
	OpenAt     string `json:"openAt"`
	CloseAt    string `json:"closeAt"`
}

type Menu struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Memo       *string `json:"memo,omitempty"`
	FileString string  `json:"fileString"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
	Rating     float32 `json:"rating"`
	IsOpen     int     `json:"isOpen"`
}

type OpenedMenu struct {
	ID              int     `json:"id"`
	MenuID          int     `json:"menuId"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Memo            *string `json:"memo,omitempty"`
	FileString      string  `json:"fileString"`
	CreateUser      string  `json:"createUser"`
	CreateDept      string  `json:"createDept"`
	OpenAt          string  `json:"openAt"`
	CloseAt         string  `json:"closeAt"`
	OrderCount      int     `json:"orderCount"`
	OrderTotalPrice int     `json:"orderTotalPrice"`
}

// Add Order Request
type Order struct {
	ID         int     `json:"id"`
	OpenMenuID int     `json:"openMenuId"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Item       string  `json:"item"`
	Sugar      int     `json:"sugar"`
	Ice        int     `json:"ice"`
	Price      int     `json:"price"`
	UserMemo   *string `json:"memo,omitempty"`
	UpdateAt   string  `json:"updateAt"`
	User       string  `json:"user"`
	Count      int     `json:"count"`
}

type Rating struct {
	ID        int     `json:"id"`
	Rating    float64 `json:"rating"`
	VoteCount int     `json:"voteCount"`
}
