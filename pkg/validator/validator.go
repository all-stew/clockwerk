package validator

import (
	"clockwerk/app/global"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	GetMessages() Messages
}

type Messages map[string]string

func GetErrorMessage(request interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		if _, isValidator := request.(Validator); isValidator {
			for _, v := range err.(validator.ValidationErrors) {
				if isValidator {
					if message, exist := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
						return message
					}
				}
			}
		}
	}
	return "参数错误"
}

func BindAndValid(c *gin.Context, v interface{}) error {
	err := c.ShouldBind(v)
	if err != nil {
		global.Log.Info(err)
		return err
	}
	return nil
}
