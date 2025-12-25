package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" form:"username" json:"username" binding:"required"`
	Password string `gorm:"not null" form:"password" json:"password" binding:"required"`
	Email    string `gorm:"unique;not null" binding:"required,email"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" binding:"required"`
	Content string `gorm:"not null" binding:"required"`
	UserID  uint   `json:"-"`
	User    User   `binding:"nostructlevel" json:"-"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User `binding:"nostructlevel" json:"-"`
	PostID  uint
	Post    Post `binding:"nostructlevel" json:"-"`
}

type PostUpdate struct {
	ID      uint   `binding:"required"`
	Title   string `binding:"required"`
	Content string `binding:"required"`
}

type CommentCreate struct {
	Content string `binding:"required"`
	PostID  uint   `binding:"required"`
}

type UserLogin struct {
	Username string `gorm:"unique;not null" form:"username" json:"username" binding:"required"`
	Password string `gorm:"not null" form:"password" json:"password" binding:"required"`
}
