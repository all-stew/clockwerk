package v1

import (
	"clockwerk/app/global"
	"clockwerk/app/views"
	"clockwerk/pkg/response"
	"github.com/gin-gonic/gin"
)

func ListUserByPage(ctx *gin.Context) {

}

func Create(ctx *gin.Context) {
	var view views.UserCreateRequestView
	err := ctx.ShouldBind(&view)

	if err != nil {
		response.NewResult(ctx).Fail(400, "参数解析失败")
		return
	}

	// todo check

	user, err := global.UserStore.Create(ctx, view.Username, view.Username, view.Phone, view.Email, view.Gender)
	if err != nil {
		response.NewResult(ctx).Fail(400, "创建失败")
		return
	}

	response.NewResult(ctx).Success(user)

}
