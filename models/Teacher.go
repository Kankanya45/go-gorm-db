// models/Teacher.go
package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model

	// Your custom fields
	FirstName string
	LastName  string
	Subjects  string
}
