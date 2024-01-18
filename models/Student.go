// models/Student.go
package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model

	// Your custom fields
	FirstName string
	LastName  string
	Age       int
	Grade     string
}
