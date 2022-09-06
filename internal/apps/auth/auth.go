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
	//type User struct {
	//	ID
	//	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	//	Mobile   string `json:"mobile" gorm:"not null;index;comment:用户手机号"`
	//	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	//	Timestamps
	//	SoftDeletes
	//}

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
