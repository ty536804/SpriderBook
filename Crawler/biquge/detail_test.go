package biquge

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestClimbDetail(t *testing.T) {
	body, err := ioutil.ReadFile("detail.html")
	if err != nil {
		panic(err)
	}
	// http://www.xbiquge.la/files/article/image/72/72327/72327s.jpg
	re := regexp.MustCompile(`<p>[\s\S]+?者：([\s\S]+?)</p>`)
	matches := re.FindSubmatch(body)
	fmt.Printf("%s \n",matches[1])
}
