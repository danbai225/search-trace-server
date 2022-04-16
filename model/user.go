package model

type User struct {
	Base
	Name  string `gorm:"column:name;type:VARCHAR(128);uniqueIndex;comment:'用户名'" json:"name"`
	Email string `gorm:"column:email;type:VARCHAR(128);uniqueIndex;comment:'邮箱'" json:"email"`
	Pass  string `gorm:"column:pass;type:VARCHAR(128);comment:'密码'" json:"pass"`
	Role  int64  `gorm:"column:role;comment:'角色'" json:"role"`
}
