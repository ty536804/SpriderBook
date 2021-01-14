package biquge

import (
	"Book/Crawler/engine"
	"regexp"
	"strconv"
)

var lastRe = regexp.MustCompile(`<a href="http://www.xbiquge.la/fenlei/([\w.]+)" class="last">([^<]+)</a>`)
func ParseLinkUl(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	totalNum,_ := strconv.Atoi(string(lastRe.FindSubmatch(contents)[2]))
	for i:=1 ;i<=totalNum; i++ {
		currentPage := strconv.Itoa(i)
		result.Request = append(result.Request,engine.Request{
			Url:      "http://www.xbiquge.la/fenlei/1_"+currentPage+".html",
			ParseFun: ParseUlList,
		})
	}
	return result
}
