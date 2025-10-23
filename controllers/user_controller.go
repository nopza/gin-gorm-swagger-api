package controllers

import (
	"net/http"
	"strconv"

	"gin-gorm-swagger-api/database"
	"gin-gorm-swagger-api/models"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get all users
// @Description Retrieve all users (JWT required)
// @Tags users
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a user (JWT required)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body models.User true "User info"
// @Success 201 {object} models.User
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// GetUserByID godoc
// @Summary Get user by ID (JWT required)
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if result := database.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update user (JWT required)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param user body models.User true "Updated info"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if result := database.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete user (JWT required)
// @Tags users
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)
	result := database.DB.Delete(&models.User{}, intID)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
