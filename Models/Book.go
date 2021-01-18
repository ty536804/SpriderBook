package Models

import (
	"encoding/json"
	"time"
)

type Book struct {
	Id               int       `json:"id" gorm:"primary_key;"`
	BookAuth         string    `json:"book_auth" gorm:"index;type:varchar(50);not null;default '';comment:'文章作者'"`
	BookTitle        string    `json:"book_title" gorm:"index;type:varchar(100);not null;default '';comment:'文章标题'"`
	BookThumbImg     string    `json:"book_thumb_img" gorm:"type:varchar(100);not null;default '';comment:'封面图' "`
	BookDesc         string    `json:"book_desc" gorm:"type:varchar(255);not null;default '';comment:'文章摘要' "`
	BookType         int       `json:"book_type" gorm:"not null;default 0;comment:'文章类型' "`
	IsHot            int       `json:"is_hot" gorm:"not null;default 0;comment:'是否热推 0否 1是' "`
	IsNum            int       `json:"is_num" gorm:"not null;default 0;comment:'点击率'"`
	IsShow           int       `json:"is_show" gorm:"not null;default 0;comment:'是否展示 0是 1否'"`
	IsFinish         int       `json:"is_finish" gorm:"not null;default 0;comment:'是否完结 0是 1否'"`
	LastArticleTitle string    `json:"last_article_title" gorm:"type:varchar(100);not null;default '';comment:'最新一篇文章标题' "`
	CreatedAt        time.Time `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt        string    `json:"updated_at" gorm:"comment:'最近一次更新时间'"`

	BookS []BookList
}

func FormJsonObj(o interface{}) (Book, error) {
	var book Book
	s, err := json.Marshal(o)
	if err != nil {
		return book, err
	}
	err = json.Unmarshal(s, &book)
	return book, err
}
