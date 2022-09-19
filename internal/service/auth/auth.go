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

type Userservice struct {
}

func (userservice Userservice) RegisterUser(params controllers.Register) (user models.User, err error) {
	var result = configs.AppObj.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
	}
	user = models.User{Name: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	err = configs.AppObj.DB.Create(&user).Error

	return
}
