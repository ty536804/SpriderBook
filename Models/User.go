package Models

type User struct {
	Id        int    `json:"id" gorm:"primary_key"`
	LoginName string `json:"login_name" gorm:"type:varchar(50);not null;default '';comment:'登陆账号' "`
	PassWd    string `json:"pass_wd" gorm:"type:char(32);not null; default '';comment:'密码'"`
}
