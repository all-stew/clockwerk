package api

import (
	"clockwerk/app/global"
	"clockwerk/pkg/response"
	"clockwerk/pkg/validator"
	"github.com/gin-gonic/gin"
)

type PingValidatorRequest struct {
	ID uint64 `form:"id" binding:"required"`
}

type PingValidatorRequest2 struct {
	Mobile   string `form:"mobile" binding:"required,is-mobile"`
	Username string `form:"username" binding:"required,is-username"`
}

func (req PingValidatorRequest2) GetMessages() validator.Messages {
	return validator.Messages{
		"Mobile.required":      "手机号不能为空",
		"Mobile.is-mobile":     "手机号校验失败",
		"Username.required":    "用户名不能为空",
		"Username.is-username": "用户名校验失败",
	}
}

func PingHandler(ctx *gin.Context) {
	response.NewResult(ctx).Success(nil)
	return
}

func PingValidatorHandler(ctx *gin.Context) {
	param := PingValidatorRequest{ID: validator.StrTo(ctx.Query("id")).MustUInt64()}
	response.NewResult(ctx).Success(param.ID)
	return
}

func PingValidatorFormHandler(ctx *gin.Context) {
	request := PingValidatorRequest{}
	err := validator.BindAndValid(ctx, &request)
	if err != nil {
		global.Log.Error(err.Error())
		response.NewResult(ctx).Fail(400, err.Error())
		return
	}
	response.NewResult(ctx).Success(request.ID)
	return
}

func PingValidatorForm2Handler(ctx *gin.Context) {
	request := PingValidatorRequest2{}
	err := validator.BindAndValid(ctx, &request)
	if err != nil {
		global.Log.Error(validator.GetErrorMessage(request, err))
		response.NewResult(ctx).Fail(400, validator.GetErrorMessage(request, err))
		return
	}
	response.NewResult(ctx).Success(request.Mobile)
	return
}
