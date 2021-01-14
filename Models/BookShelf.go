package Models

type BookShelf struct {
	Id     int `json:"id" gorm:"primary_key;comment:'ID'"`
	UserId int `json:"user_id" gorm:"index;not null;default 0;comment:'用户ID'"`
	BookId int `json:"book_id" gorm:"index;not null;default 0;comment:'书ID' "`
}
