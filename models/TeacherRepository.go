// models/TeacherRepository.go
package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	Db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) *TeacherRepository {
	return &TeacherRepository{Db: db}
}

func (r *TeacherRepository) GetTeachers(c *gin.Context) {
	var teachers []Teacher
	r.Db.Find(&teachers)
	c.JSON(200, teachers)
}

func (r *TeacherRepository) GetTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher Teacher
	r.Db.First(&teacher, id)
	c.JSON(200, teacher)
}

func (r *TeacherRepository) CreateTeacher(c *gin.Context) {
	var newTeacher Teacher
	c.BindJSON(&newTeacher)
	r.Db.Create(&newTeacher)
	c.JSON(200, newTeacher)
}

func (r *TeacherRepository) UpdateTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher Teacher
	r.Db.First(&teacher, id)
	c.BindJSON(&teacher)
	r.Db.Save(&teacher)
	c.JSON(200, teacher)
}

func (r *TeacherRepository) DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher Teacher
	r.Db.Delete(&teacher, id)
	c.JSON(200, gin.H{"id" + id: "is deleted"})
}
