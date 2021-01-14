package Book

import (
	"Book/Pkg/Helpers"
	"Book/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Index(c *gin.Context) {
	data := make(map[string]interface{})
	book := Services.NewBookService()
	page := com.StrTo(c.Query("page")).MustInt()
	pageNum := Helpers.PageNum(com.StrTo(c.Query("pageNum")).MustInt())
	data["list"] = book.GetBooks(page, pageNum)
	data["count"] = 0
	data["size"] = pageNum
	Helpers.SuccessRes(c, Helpers.SUCCESS, "导航", data)
}
