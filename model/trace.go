package model

type Trace struct {
	Base
	Title   string `gorm:"column:title;type:VARCHAR(128);comment:'标题'" json:"title"`
	Url     string `gorm:"column:url;type:TEXT;comment:'超链接'" json:"url"`
	Content string `gorm:"column:content;type:LONGTEXT;comment:'内容文本'" json:"content"`
}
