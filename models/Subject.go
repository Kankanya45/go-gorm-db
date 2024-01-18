// models/Subject.go
package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model

	// Your custom fields
	Name        string
	Description string
}
