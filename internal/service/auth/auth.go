package auth

import (
	"errors"
	"gitee.com/bytesworld/tomato/configs"
	. "gitee.com/bytesworld/tomato/internal"
	"gitee.com/bytesworld/tomato/internal/controllers"
	"gitee.com/bytesworld/tomato/internal/models"
	"gitee.com/bytesworld/tomato/pkg/utils"
)

func GetUserByID(id string) (user models.User, err error) {
	err = DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

type userservice struct {
}

var UserService =new(userservice)

func (userservice userservice) RegisterUser(params controllers.Register) (user models.User, err error) {
	var result = configs.AppObj.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
	}
	user = models.User{Name: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	err = configs.AppObj.DB.Create(&user).Error

	return
}
func (userservice *userservice) Login(params controllers.Login) (err error, user *models.User) {
	err = configs.AppObj.DB.Where("mobile=?",params.Mobile).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}