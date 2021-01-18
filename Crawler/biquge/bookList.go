package biquge

import (
	"Book/Crawler/engine"
	"Book/Pkg/Helpers"
	"Book/Services"
	"fmt"
	"regexp"
	"strings"
)

var bookListRe = regexp.MustCompile(`<dd><a href='([a-zA-Z0-9/.]+)'[^>]*>([^<]+)</a></dd>`)

var imgRe = regexp.MustCompile(`<img alt="(.+)" src="(http://www.xbiquge.la/files/article/[a-zA-Z0-9./]+)" [^>]*>`)
var upTimeRe = regexp.MustCompile(`<p>最后更新：([^<]+)</p>`)
var nearSectionRe = regexp.MustCompile(`<p>最新章节：<a href="http://www.xbiquge.la/[a-zA-Z0-9./]+">([^<]+)</a></p>`)
var autoRe = regexp.MustCompile(`<p>[\s\S]+?者：([\s\S]+?)</p>`)
var descRe = regexp.MustCompile(`<div id="intro">([\s\S]+?)</div>`)

// 小书列表页
func ParseBookList(contents []byte, uList string) engine.ParseResult {
	result := engine.ParseResult{}
	if len(contents) < 1 {
		fmt.Println("空的内容，小书:")
		return result
	}
	//matches := bookListRe.FindAllSubmatch(contents,-1)
	item := make(map[string]interface{})
	bookImg := imgRe.FindSubmatch(contents)
	item["book_thumb_img"] = string(bookImg[2])

	CreatedAt := upTimeRe.FindSubmatch(contents)
	item["updated_at"] = string(CreatedAt[1])

	item["book_title"] = string(bookImg[1])
	nearSection := nearSectionRe.FindSubmatch(contents)
	item["last_article_title"] = string(nearSection[1])
	isFinish := 0
	if strings.Index(string(nearSection[1]), "完结") > 0 {
		isFinish = 1
	}
	item["is_finish"] = isFinish
	item["book_type"] = 1
	item["is_hot"] = 1
	item["is_num"] = 1
	item["is_show"] = 0
	desc := descRe.FindSubmatch(contents)
	conRe := regexp.MustCompile(`<p>([^<font])([\s\S]+)</p>`)
	cc := conRe.FindSubmatch(desc[1])
	item["book_desc"] = string(cc[0])

	bookAuto := autoRe.FindSubmatch(contents)
	item["book_auth"] = string(bookAuto[1])
	book := Services.NewBookService()

	code, msg, bookId := book.AddBook(item)
	if code == Helpers.SUCCESS {
		result.Items = append(result.Items, engine.Item{
			Url:     uList,
			Type:    "玄幻",
			Id:      string(bookId),
			Payload: item,
		})
		//for _, m := range matches{
		//	bookName := string(m[2])
		//	result.Request = append(result.Request,engine.Request{
		//		Url:      "http://www.xbiquge.la"+string(m[1]),
		//		ParseFun: func(bytes []byte) engine.ParseResult {
		//			return ClimbDetail(bytes,bookId,bookName)
		//		},
		//	})
		//}
	} else {
		fmt.Println("文章添加失败:", msg)
	}
	return result
}
