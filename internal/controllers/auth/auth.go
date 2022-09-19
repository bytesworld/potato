package auth

import (
	. "gitee.com/bytesworld/tomato/internal"
	"gitee.com/bytesworld/tomato/internal/controllers"
	. "gitee.com/bytesworld/tomato/internal/logger"
	"gitee.com/bytesworld/tomato/internal/models"
	sv_auth "gitee.com/bytesworld/tomato/internal/service/auth"
	"gitee.com/bytesworld/tomato/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	Logger.Info("login")
	c.JSON(http.StatusOK, gin.H{
		"message": "auth",
	})
}

func Register(c *gin.Context) {
	Logger.Info("register user")
	var form controllers.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		Logger.Error(err)
		return
	}

	user, err := sv_auth.Userservice{}.RegisterUser(form)
	if err != nil {
		Logger.Error(err)
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

func GetUsers(c *gin.Context) {
	var users []models.User
	DB.Find(&users)
	c.JSON(http.StatusCreated, users)
}
func CreateUser(c *gin.Context) {
	//var _user models.CreateUser
	var _user controllers.Register
	if err := c.ShouldBindJSON(&_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": controllers.GetErrorMsg(_user, err)})
		return
	}
	user := models.User{Name: _user.Name, Mobile: _user.Mobile, Password: _user.Password}
	DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	user, err := sv_auth.GetUserByID(c.Param("id"))
	if err != nil {
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
