package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/entity"
	"github.com/pai000/sa-65-project/service"
	"golang.org/x/crypto/bcrypt"
)

// LoginPayload login body
type LoginPayload struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignUpPayload signup body
type SignUpPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginResponse struct {
	Position string `json:"position"`
	Token    string `json:"token"`
	ID       uint   `json:"id"`
}

// POST /login
func LoginEmployee(c *gin.Context) {
	var payload LoginPayload
	var em entity.Employee

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา user ด้วย email ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM employees WHERE email = ?", payload.Email).Scan(&em).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(em.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is incerrect"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(em.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginResponse{
		Position: GetPositionName(em.ID),
		Token:    signedToken,
		ID:       em.ID,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}

func LoginStudent(c *gin.Context) {
	var payload LoginPayload
	var std entity.Student

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา user ด้วย email ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM students WHERE STUDENT_NUMBER = ?", payload.Email).Scan(&std).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(std.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is incerrect"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(std.STUDENT_NUMBER)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginResponse{
		Position: GetRoleName(std.ID),
		Token:    signedToken,
		ID:       std.ID,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}

func GetPositionName(id uint) string {
	rn := entity.Employee{}
	tx := entity.DB().Preload("Job_Position").First(&rn, id)

	if tx.Error != nil {
		return "Role not found"
	} else if rn.Job_Position.Name == "Admin" {
		return "Admin"
	}
	return "err"
}

func GetRoleName(id uint) string {
	rn := entity.Student{}
	tx := entity.DB().Preload("Role").First(&rn, id)

	if tx.Error != nil {
		return "Role not found"
	} else if rn.Role.Role_name == "Student" {
		return "Student"
	}
	return "err"
}

// POST /create employee
func CreateLoginEmployee(c *gin.Context) {
	var payload SignUpPayload
	var em entity.Employee

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}

	em.Name = payload.Name
	em.Email = payload.Email
	em.Personal_ID = string(hashPassword)

	if err := entity.DB().Create(&em).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": em})
}

// POST /create employee
func CreateLoginStudent(c *gin.Context) {
	var payload SignUpPayload
	var em entity.Student

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}

	em.STUDENT_NAME = payload.Name
	em.STUDENT_NUMBER = payload.Username
	em.PERSONAL_ID = string(hashPassword)

	if err := entity.DB().Create(&em).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": em})
}
