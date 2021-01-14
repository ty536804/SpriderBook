package biquge

import (
	"Book/Crawler/engine"
	"Book/Pkg/Helpers"
	"Book/Services"
	"fmt"
	"regexp"
)

// @Summer 获取详情
var conRe = regexp.MustCompile(`<div id="content">([\w\S\s]+?)</div>`)
func ClimbDetail(contents []byte, bookId int,bookName string) engine.ParseResult {
	data := make(map[string]interface{})
	if len(contents) < 1 {
		fmt.Println("小书详情获取失败:",bookName)
		return engine.ParseResult{}
	}
	articleContent := conRe.FindSubmatch(contents)
	data["book_id"] = bookId
	data["article_title"] = bookName
	data["article_content"] = string(articleContent[1])
	bookList := Services.NewBookListService()
	code, msg := bookList.AddBookList(data)
	if code != Helpers.SUCCESS {
		fmt.Println(msg)
	}
	return engine.ParseResult{}
}
