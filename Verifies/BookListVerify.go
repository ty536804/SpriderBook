package Verifies

import (
	"Book/Pkg/Helpers"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type bookListVerify interface {
	VerifyBookList(c *gin.Context) (code int, err string, res map[string]interface{})
}

type bookListVerifies struct {
}

func NewBookListVerify() *bookListVerifies {
	return &bookListVerifies{}
}

func (b *bookListVerifies) VerifyBookList(data map[string]interface{}) (code int, err string) {
	BookId := data["book_id"]
	ArticleTitle := data["article_title"]
	ArticleContent := data["article_content"]

	valid := validation.Validation{}
	valid.Required(BookId, "book_id").Message("书ID")
	valid.Required(ArticleTitle, "article_title").Message("文章标题不能为空")
	valid.Required(ArticleContent, "article_content").Message("正文不能为空")

	if !valid.HasErrors() {
		return Helpers.SUCCESS, "操作成功"
	}
	return Helpers.ViewErr(valid)
}
