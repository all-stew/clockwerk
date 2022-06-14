package api

import (
	"clockwerk/app/global"
	"clockwerk/pkg/response"
	"clockwerk/pkg/validator"
	"github.com/gin-gonic/gin"
)

type PingValidatorRequest struct {
	ID uint64 `form:"id" binding:"required,min=1,max=20"`
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
	valid, err := validator.BindAndValid(ctx, &request)
	if !valid {
		global.Log.Error(err.Error())
		response.NewResult(ctx).Fail(400, err.Error())
		return
	}
	response.NewResult(ctx).Success(request.ID)
	return
}
