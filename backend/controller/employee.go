package controller

import (
	"github.com/pai000/sa-65-project/entity"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users
func CreateEmployee(c *gin.Context) {

	var employee entity.Employee
	var gender entity.Gender
	var job_position entity.Job_Position
	var province entity.Province

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 7 จะถูก bind เข้าตัวแปร employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 8: ค้นหา gender ด้วย id
	if tx := entity.DB().Where("id = ?", employee.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// 9: ค้นหา job_position ด้วย id
	if tx := entity.DB().Where("id = ?", employee.Job_PositionID).First(&job_position); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "job_position not found"})
		return
	}

	// 10: ค้นหา province ด้วย id
	if tx := entity.DB().Where("id = ?", employee.ProvinceID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}
	// เข้ารหัสลับจากบัตรประชาชนที่ Admin กรอกข้อมูล
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(employee.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}

	employee.Password = string(hashPassword)

	// 12: สร้าง Employee
	wv := entity.Employee{
		Name:         employee.Name,
		Personal_ID:  employee.Personal_ID,
		Email:        employee.Email,
		Password:     employee.Password,
		Gender:       gender,       // โยงความสัมพันธ์กับ Entity gender
		Job_Position: job_position, // โยงความสัมพันธ์กับ Entity job_position
		Province:     province,     // โยงความสัมพันธ์กับ Entity province
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /employee/:id
func GetEmployee(c *gin.Context) {
	var employee entity.Employee
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// GET /employees

func ListEmployee(c *gin.Context) {

	var employees []entity.Employee

	if err := entity.DB().Preload("Gender").Preload("Job_Position").Preload("Province").Raw("SELECT * FROM employees").Find(&employees).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": employees})

}

// DELETE /employees/:id

func DeleteEmployee(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM employees WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /employees

func UpdateEmployee(c *gin.Context) {

	var employee entity.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", employee.ID).First(&employee); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})

		return

	}

	if err := entity.DB().Save(&employee).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": employee})

}
