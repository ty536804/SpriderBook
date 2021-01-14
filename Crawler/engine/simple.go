package engine

import (
	"Book/Crawler/Fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := Fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error "+"fetching url:%s:%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFun(body), nil
}
