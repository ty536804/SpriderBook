package biquge

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestParseBookList(t *testing.T) {
	body, err := ioutil.ReadFile("list.html")
	if err != nil {
		fmt.Print("获取文件失败")
		panic(err)
	}

	re := regexp.MustCompile(`《<a href="(http://www.xbiquge.la/[a-zA-Z0-9/]+)" [^>]*>([^<]+)</a>》</span>`)
	matches :=  re.FindAllSubmatch(body,-1)
	for _, m := range matches{

			fmt.Printf("Url: %s : bookName: %s",m[1],m[2])

	}
	//fmt.Printf("%s",matches)
	fmt.Println()
}
