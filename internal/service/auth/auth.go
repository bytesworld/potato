package auth

import (
	. "gitee.com/bytesworld/tomato/internal"
	"gitee.com/bytesworld/tomato/internal/models"
)

func GetUserByID(id string) (user models.User, err error) {
	err = DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
