package biquge

import (
	"Book/Crawler/engine"
	"regexp"
)

// 解析列表
func ParseUlList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(`《<a href="(http://www.xbiquge.la/[a-zA-Z0-9/]+)" [^>]*>([^<]+)</a>》`)
	matches :=  re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _, m := range matches{
		bookName := string(m[2])
		result.Request = append(result.Request,engine.Request{
			Url:      string(m[1]),
			ParseFun: func(bytes []byte) engine.ParseResult {
				return ParseBookList(bytes,bookName)
			},
		})
		result.Items = append(result.Items,bookName)
	}
	return result
}
