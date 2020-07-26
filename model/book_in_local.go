package model

import (
	"github.com/jinzhu/gorm"
)

type TypeInfo struct {
	Type string
}

type BookInLocal struct {
	gorm.Model
	Title string
	Type  string
}
