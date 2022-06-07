package sys

import (
	"clockwerk/pkg/response"
	"clockwerk/src/global"
	"clockwerk/src/views"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var view views.UserCreateRequestView
	err := ctx.ShouldBind(&view)

	if err != nil {
		response.Fail(ctx, 200, 400, "参数解析失败", nil)
		return
	}

	// todo check

	user, err := global.UserStore.Create(ctx, view.Username, view.Username, view.Phone, view.Email, view.Gender)
	if err != nil {
		response.Fail(ctx, 200, 400, "创建失败", nil)
		return
	}

	if err != nil {
		response.Fail(ctx, 200, 200, "创建成功", user)
		return
	}

}
