package entity

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Role_name string
	Students  []Student `gorm:"foreignKey:RoleID"`
}
type Program struct {
	gorm.Model
	Program_name string
	Students     []Student `gorm:"foreignKey:ProgramID"`
}
type Student struct {
	gorm.Model
	STUDENT_NUMBER string
	STUDENT_NAME   string
	PERSONAL_ID    string
	Password       string

	//
	GenderID *uint
	Gender   Gender
	//
	ProvinceID *uint
	Province   Province
	//
	ProgramID *uint
	Program   Program
	//
	RoleID *uint
	Role   Role
	//
	EmployeeID *uint
	Employee   Employee
}
