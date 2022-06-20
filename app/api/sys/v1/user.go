package v1

import (
	"clockwerk/app/service"
	"clockwerk/app/views"
	"clockwerk/pkg/response"

	"github.com/gin-gonic/gin"
)

type userAPIHandler struct {
	userService service.UserService
}

func NewUserAPIHandler(u service.UserService) *userAPIHandler {
	return &userAPIHandler{
		userService: u,
	}
}

func ListUserByPage(ctx *gin.Context) {

}

func (u *userAPIHandler) Create(ctx *gin.Context) {
	var view views.UserCreateRequestView
	err := ctx.ShouldBind(&view)

	if err != nil {
		response.NewResult(ctx).Fail(400, "参数解析失败")
		return
	}

	user, err := u.userService.Create(view.Nickname, view.Phone, view.Email, view.Gender)
	if err != nil {
		response.NewResult(ctx).Fail(400, "创建失败")
		return
	}

	response.NewResult(ctx).Success(user)

}
