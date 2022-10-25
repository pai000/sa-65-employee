package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/entity"
)

// POST /job_position
func CreateJob_Position(c *gin.Context) {
	var job_position entity.Job_Position
	if err := c.ShouldBindJSON(&job_position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&job_position).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": job_position})
}

// GET /job_position
// List all job_position
func ListJob_Position(c *gin.Context) {
	var job_positions []entity.Job_Position
	if err := entity.DB().Raw("SELECT * FROM job_positions").Scan(&job_positions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": job_positions})
}

// GET /job_position/:id
// Get job_position by id
func GetJob_Position(c *gin.Context) {
	var job_positions entity.Job_Position
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&job_positions); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "job_position not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": job_positions})
}

// DELETE /job_position/:id
func DeleteJob_Position(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Job_Positions WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "job_position not found"})
		return
	}
}

// PATCH /job_positions
func UpdateJob_Position(c *gin.Context) {
	var job_position entity.Job_Position
	if err := c.ShouldBindJSON(&job_position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", job_position.ID).First(&job_position); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "job_position not found"})
		return
	}

	if err := entity.DB().Save(&job_position).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": job_position})
}
