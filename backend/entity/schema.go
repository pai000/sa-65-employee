package entity

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	Name string

	Employees []Employee `gorm:"foreignKey:GenderID"`
	Student   []Student  `gorm:"foreignKey:GenderID"`
}

type Job_Position struct {
	gorm.Model
	Name string

	Employees []Employee `gorm:"foreignKey:Job_PositionID"`
	Student   []Student  `gorm:"foreignKey:GenderID"`
}

type Province struct {
	gorm.Model
	Name string

	Employees []Employee `gorm:"foreignKey:ProvinceID"`
}

type Employee struct {
	gorm.Model
	Personal_ID string `gorm:"uniqueIndex"`
	Email       string `gorm:"uniqueIndex"`
	Name        string
	Password    string

	//GenderID ทำหน้าที่เป็น FK
	GenderID *uint
	Gender   Gender `gorm:"references:id"`

	//Job_PositionID ทำหน้าที่เป็น FK
	Job_PositionID *uint
	Job_Position   Job_Position `gorm:"references:id"`

	//ProvinceID ทำหน้าที่เป็น FK
	ProvinceID *uint
	Province   Province `gorm:"references:id"`

	RoleID *uint
	Role   Role `gorm:"references:id"`
}
