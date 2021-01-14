package Services

import (
	db "Book/Database"
	"Book/Models"
	"Book/Pkg/Helpers"
	"Book/Verifies"
	"fmt"
)

type IOBookService interface {
	AddBook(data map[string]interface{}) (code int, msg string, id int)
	GetBook(auth, title string) (book Models.Book)
	GetBookById(id string) (book Models.Book)
	GetBooks(page, pageNum int) (book []Models.Book)
	GetHotBook(pageNum int) (book []Models.Book)
}

type bookService struct {
}

func NewBookService() *bookService {
	return &bookService{}
}

func (b *bookService) AddBook(data map[string]interface{}) (code int, msg string, id int) {
	bookVerify := Verifies.NewBookVerify()
	code, msg = bookVerify.VerifyBook(data)
	if code == Helpers.ERROR {
		return code, msg, 0
	}
	return addBook(data)
}

// @Summer 添加文章
func addBook(data map[string]interface{}) (code int, err string, id int) {
	user := &Models.Book{
		BookAuth:         data["book_auth"].(string),
		BookTitle:        data["book_title"].(string),
		BookThumbImg:     data["book_thumb_img"].(string),
		BookDesc:         data["book_desc"].(string),
		BookType:         data["book_type"].(int),
		IsHot:            data["is_hot"].(int),
		IsNum:            data["is_num"].(int),
		IsShow:           data["is_show"].(int),
		IsFinish:         data["is_finish"].(int),
		LastArticleTitle: data["last_article_title"].(string),
		UpdatedAt:        data["updated_at"].(string),
	}
	res := db.Db.Create(user)
	if res.Error != nil {
		fmt.Println("文章入库失败:", res)
		return Helpers.ERROR, res.Error.Error(), 0
	}
	return Helpers.SUCCESS, "", user.Id
}

// @Summer通过作者名称或书名称查找
func (b *bookService) GetBook(auth, title string) (book Models.Book) {
	db.Db.Select("book_auth,book_title,last_article_title,updated_at").
		Where("book_auth like ? ,book_title like ?").Find(&book)
	return
}

// @Summer通过书ID获取书所有章节
func (b *bookService) GetBookById(id string) (book Models.Book) {
	db.Db.Where("id = ? ", id).Related(&Models.BookList{}).Find(&book)
	return
}

// 分页数据
func (b *bookService) GetBooks(page, pageNum int) (book []Models.Book) {
	offset := Helpers.PageNum(page)
	limit := Helpers.Offset(page, pageNum)
	db.Db.Order("updated_at desc").Limit(limit).Offset(offset).Find(&book)
	return
}

// @Summer 获取点击率最高，一定数量的文章
func (b *bookService) GetHotBook(pageNum int) (book []Models.Book) {
	db.Db.Limit(pageNum).Order("is_num desc,is_hot desc").Where("is_show = 0").Find(&book)
	return
}

// @Summer 删除表
func DropTable() {
	if db.Db.HasTable(&Models.Book{}) {
		db.Db.DropTable(&Models.Book{})
	}
}

// @Summer 生成表
func MigrateTable() {
	if !db.Db.HasTable(&Models.Book{}) {
		db.Db.AutoMigrate(&Models.Book{})
	}
}

// @Summer 初始化表
func InitTable() {
	DropTable()
	db.Db.AutoMigrate(&Models.Book{})
}
