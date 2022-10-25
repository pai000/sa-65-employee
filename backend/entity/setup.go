package entity

import (
	//"fmt"
	//"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Gender{},
		&Job_Position{},
		&Province{},
		&Employee{},
	)

	db = database

	//add example data

	//Gender
	gender1 := Gender{
		Name: "Male",
	}

	db.Model(&Gender{}).Create(&gender1)

	gender2 := Gender{
		Name: "FeMale",
	}

	db.Model(&Gender{}).Create(&gender2)

	//insert job_position
	job_position1 := Job_Position{
		Name: "Admin",
	}
	db.Model(&Job_Position{}).Create(&job_position1)

	job_position2 := Job_Position{
		Name: "Housekeeper",
	}
	db.Model(&Job_Position{}).Create(&job_position2)

	job_position3 := Job_Position{
		Name: "Security Guard",
	}
	db.Model(&Job_Position{}).Create(&job_position3)

	job_position4 := Job_Position{
		Name: "Mechanic",
	}
	db.Model(&Job_Position{}).Create(&job_position4)

	//province
	korat := Province{
		Name: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&korat)

	chon := Province{
		Name: "Chonburi",
	}
	db.Model(&Province{}).Create(&chon)

	bangkok := Province{
		Name: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	password1, err := bcrypt.GenerateFromPassword([]byte("abc12456"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)

	//insert employee
	db.Model(&Employee{}).Create(&Employee{
		Personal_ID: string(password1),
		Email:       "ana@gmail.com",
		Name:        "Ana poul",

		Gender:       gender2,
		Job_Position: job_position1,
		Province:     korat,
	})

	db.Model(&Employee{}).Create(&Employee{
		Personal_ID: string(password3),
		Email:       "kerkkiat@gmail.com",
		Name:        "Kerkkiat Prabmontree",

		Gender:       gender1,
		Job_Position: job_position3,
		Province:     bangkok,
	})

	db.Model(&Employee{}).Create(&Employee{
		Personal_ID: string(password2),
		Email:       "matinez@gmail.com",
		Name:        "Devid Matinez",

		Gender:       gender1,
		Job_Position: job_position4,
		Province:     chon,
	})

}
