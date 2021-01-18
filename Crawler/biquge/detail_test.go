package biquge

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestClimbDetail(t *testing.T) {
	file, err := ioutil.ReadFile("list.html")

	//resp, err := http.Get("http://www.xbiquge.la/10/10489/")
	if err != nil {
		panic(err)
	}
	//defer resp.Body.Close()
	//res, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	var re = regexp.MustCompile(`<img alt="(.+)" src="(http://www.xbiquge.la/files/article/[a-zA-Z0-9./]+)" [^>]*>`)
	matches := re.FindSubmatch(file)
	fmt.Printf("%s \n", matches[1])
}
