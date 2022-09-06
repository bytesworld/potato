package auth

import (
	. "gitee.com/bytesworld/tomato/internal"
	. "gitee.com/bytesworld/tomato/internal/logger"
	"gitee.com/bytesworld/tomato/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	Logger.Info("login")
	c.JSON(http.StatusOK, gin.H{
		"message": "auth",
	})
}

func AddUser(c *gin.Context) {
	Logger.Info("Add user")
	DB.Create(
		&models.User{
			Name:     "weidong",
			Mobile:   "19211111",
			Password: "123",
		})
	var user models.User
	DB.First(&user, 1)
	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	DB.Find(&users)
	c.JSON(http.StatusCreated, users)
}
func CreateUser(c *gin.Context) {
	var _user models.CreateUser
	if err := c.ShouldBindJSON(&_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Name: _user.Name, Mobile: _user.Mobile, Password: _user.Password}
	DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	var user models.User

	if err := DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, user)
}
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var _user models.CreateUser
	if err := c.ShouldBindJSON(&_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&user).Updates(_user)

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	DB.Delete(&user)

	c.JSON(http.StatusOK, user)
}
