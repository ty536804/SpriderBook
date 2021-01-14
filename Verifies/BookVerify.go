package Verifies

import (
	"Book/Pkg/Helpers"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type bookVerify interface {
	VerifyAdmin(c *gin.Context) (code int, err string, res map[string]interface{})
}

type bookVerifies struct {
}

func NewBookVerify() *bookVerifies {
	return &bookVerifies{}
}

func (a *bookVerifies) VerifyBook(data map[string]interface{}) (code int, err string) {
	BookAuth := data["book_auth"]
	BookTitle := data["book_title"]
	BookThumbImg := data["book_thumb_img"]
	BookDesc := data["book_desc"]
	BookType := data["book_type"]

	valid := validation.Validation{}
	valid.Required(BookAuth, "book_auth").Message("文章作者不能为空")
	valid.Required(BookTitle, "book_title").Message("文章标题不能为空")
	valid.Required(BookThumbImg, "book_thumb_img").Message("封面图不能为空")
	valid.Required(BookDesc, "book_desc").Message("文章摘要不能为空")
	valid.Required(BookType, "book_type").Message("文章类型不能为空")

	if !valid.HasErrors() {
		return Helpers.SUCCESS, "操作成功"
	}
	return Helpers.ViewErr(valid)
}
