package biquge

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
)

func TestParseBookList(t *testing.T) {
	resp, err := http.Get("http://www.xbiquge.la/")
	if err != nil {
		fmt.Print("获取文件失败")
		panic(err)
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`《<a href="(http://www.xbiquge.la/[a-zA-Z0-9/]+)" [^>]*>([^<]+)</a>》</span>`)
	matches := re.FindAllSubmatch(res, -1)
	for _, m := range matches {

		fmt.Printf("Url: %s : bookName: %s", m[1], m[2])

	}
	//fmt.Printf("%s",matches)
	fmt.Println()
}
