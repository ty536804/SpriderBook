package Services

import (
	db "Book/Database"
	"Book/Models"
	"Book/Pkg/Helpers"
	"Book/Verifies"
	"fmt"
)

type IOBookListService interface {
	AddBookList(data map[string]interface{}) (code int, msg string)
	GetBookListById(id string) (book Models.BookList)
	GetBookLists(page, pageNum int) (book []Models.BookList)
}

type bookListService struct {
}

func NewBookListService() *bookListService {
	return &bookListService{}
}

func (b *bookListService) AddBookList(data map[string]interface{}) (code int, msg string) {
	bookVerify := Verifies.NewBookListVerify()
	code, msg = bookVerify.VerifyBookList(data)
	if code == Helpers.ERROR {
		return code, msg
	}
	return addBookList(data)
}

// @Summer 添加文章
func addBookList(data map[string]interface{}) (code int, err string) {
	res := db.Db.Create(&Models.BookList{
		BookId:         data["book_id"].(int),
		ArticleTitle:   data["article_title"].(string),
		ArticleContent: data["article_content"].(string),
	})
	if res.Error != nil {
		fmt.Println("文章入库失败")
		return Helpers.Error()
	}
	return Helpers.Success()
}

// @Summer通过书ID获取书所有章节
func (b *bookListService) GetBookListById(id string) (book Models.BookList) {
	db.Db.Where("id = ? ", id).Related(&Models.BookList{}).Find(&book)
	return
}

// 分页数据
func (b *bookListService) GetBookLists(page, pageNum int) (book []Models.BookList) {
	offset := Helpers.PageNum(page)
	limit := Helpers.Offset(page, pageNum)
	db.Db.Order("updated_at desc").Limit(limit).Offset(offset).Find(&book)
	return
}

// @Summer 删除表
func DropBookListTable() {
	if db.Db.HasTable(&Models.BookList{}) {
		db.Db.DropTable(&Models.BookList{})
	}
}

// @Summer 生成表
func MigrateBookListTable() {
	if !db.Db.HasTable(&Models.BookList{}) {
		db.Db.AutoMigrate(&Models.BookList{})
	}
}

// @Summer 初始化表
func InitBookListTable() {
	DropBookListTable()
	db.Db.AutoMigrate(&Models.BookList{})
}
