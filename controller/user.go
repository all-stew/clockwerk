package controller

import (
	"clockwerk/dao"
	"clockwerk/entity/models"
	"clockwerk/entity/views"
	"clockwerk/entity/views/response"
	"clockwerk/utils"
	"github.com/gin-gonic/gin"
)

// GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	// 获取参数
	UserListForm := views.UserListForm{}
	if err := c.ShouldBind(&UserListForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 获取数据
	total, userlist := dao.GetUserListDao(UserListForm.Page, UserListForm.PageSize)
	// 判断
	if (total + len(userlist)) == 0 {
		response.Fail(c, 400, 400, "未获取到到数据", map[string]interface{}{
			"total":    total,
			"userlist": userlist,
		})
		return
	}
	response.Success(c, 200, "获取用户列表成功", map[string]interface{}{
		"total":    total,
		"userlist": userlist,
	})
}

// PasswordLogin 登录
func PasswordLogin(c *gin.Context) {
	PasswordLoginForm := views.PasswordLoginForm{}
	if err := c.ShouldBind(&PasswordLoginForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	//// 数字验证码验证失败    store.Verify(验证码id,验证码,验证后是否关闭)
	//if !store.Verify(PasswordLoginForm.CaptchaId, PasswordLoginForm.Captcha, true) {
	//	Response.Err(c, 400, 400, "验证码错误", "")
	//	return
	//}
	//查询数据库是否有改用户
	user, ok := dao.FindUserInfo(PasswordLoginForm.Username, PasswordLoginForm.PassWord)
	if !ok {
		response.Fail(c, 401, 401, "未注册该用户", "")
		return
	}
	//
	token := utils.CreateToken(c, user.ID, user.NickName, user.Role)
	userinfoMap := HandleUserModelToMap(user)
	userinfoMap["token"] = token
	response.Success(c, 200, "success", userinfoMap)
}

func HandleUserModelToMap(user *models.User) map[string]interface{} {
	birthday := ""
	if user.Birthday == nil {
		birthday = ""
	} else {
		birthday = user.Birthday.Format("2006-01-02")
	}
	userItemMap := map[string]interface{}{
		"id":        user.ID,
		"nick_name": user.NickName,
		"head_url":  user.HeadUrl,
		"birthday":  birthday,
		"address":   user.Address,
		"desc":      user.Desc,
		"gender":    user.Gender,
		"role":      user.Role,
		"mobile":    user.Mobile,
	}
	return userItemMap
}
