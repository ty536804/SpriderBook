package engine

type BookMenu struct {
	BookAuth     string `json:"book_auth"`
	BookTitle    string `json:"book_title"`
	BookThumbImg string `json:"book_thumb_img"`
	BookDesc     string `json:"book_desc"`
}

type Request struct {
	Url      string
	ParseFun func([]byte) ParseResult
}

type ParseResult struct {
	Request []Request
	Items   []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParseFun([]byte) ParseResult {
	return ParseResult{}
}
