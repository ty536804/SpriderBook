package biquge

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
)

func TestClimbDetail(t *testing.T) {
	resp, err := http.Get("http://www.xbiquge.la/10/10489/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`<p>[\s\S]+?者：([\s\S]+?)</p>`)
	matches := re.FindSubmatch(res)
	fmt.Printf("%s \n", matches[1])
}
