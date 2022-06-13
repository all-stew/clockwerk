package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Result struct {
	Ctx *gin.Context
}

type ApiResponse struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Time    int64       `json:"time"`
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

func (r *Result) Success(data interface{}) {
	res := ApiResponse{}
	res.Code = 200
	res.Message = "success"
	res.Data = data
	res.Time = time.Now().Unix()
	r.Ctx.JSON(http.StatusOK, res)
}

func (r *Result) Fail(code uint, msg string) {
	res := ApiResponse{}
	res.Code = code
	res.Message = msg
	res.Time = time.Now().Unix()
	r.Ctx.JSON(http.StatusOK, res)
}
