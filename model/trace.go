package model

type Trace struct {
	ID       int64  `gorm:"column:id;type:BIGINT;autoIncrement;comment:'唯一id';NOT NULL" json:"id"`
	UserName string `gorm:"column:username;type:VARCHAR(128);index;comment:'用户名'" json:"username"`
	Title    string `gorm:"column:title;type:VARCHAR(128);comment:'标题'" json:"title"`
	Url      string `gorm:"column:url;type:TEXT;comment:'超链接'" json:"url"`
	Content  string `gorm:"column:content;type:LONGTEXT;comment:'内容文本'" json:"content"`
	Base
}
