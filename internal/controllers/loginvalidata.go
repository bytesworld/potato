package controllers

type Login struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required" `
	Password string `json:"password" form:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Mobile.required":   "手机号码不能为空",
		//"Mobile.mobile": "手机号码格式不正确",
		"Password.required": "用户密码不能为空",
	}
}