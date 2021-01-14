package Models

type BookList struct {
	Id             int    `json:"id" gorm:"primary_key"`
	BookId         int    `json:"book_id" gorm:"index;not null;default 0;comment:'书ID' "`
	ArticleTitle   string `json:"article_title" gorm:"type:varchar(100);not null;default '';comment:'文章标题' "`
	ArticleContent string `json:"article_content" gorm:"type:text;not null;default ''; comment:'正文' "`
}
