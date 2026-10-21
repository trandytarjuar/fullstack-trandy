package main

type User struct {
	UserID    uint   `json:"userId" gorm:"primaryKey;column:user_id"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
}

type UserRequest struct {
	FirstName string `json:"firstName" binding:"required,max=100"`
	LastName  string `json:"lastName" binding:"required,max=100"`
	Email     string `json:"email" binding:"required,email,max=254"`
}
