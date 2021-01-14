package Helpers

import (
	"Book/Pkg/Setting"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// @Summer 解析错误原因
func ViewErr(valid validation.Validation) (code int, err string) {
	for _, err := range valid.Errors {
		return ERROR, err.Message
	}
	return SUCCESS, "操作成功"
}

func Error() (code int, err string) {
	return ERROR, "操作失败"
}

func Success() (code int, err string) {
	return SUCCESS, "操作成功"
}

// @Summary 返回正确内容
func SuccessRes(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.SecureJSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// @Summer 每页显示数量
func PageNum(pageNum int) int {
	if pageNum < 1 || pageNum > Setting.MaxPageSize {
		return Setting.PageSize
	}
	return pageNum
}

// @Summer  分页偏移量
func Offset(page, pageNum int) int {
	offset := (page - 1) * pageNum
	return offset
}
