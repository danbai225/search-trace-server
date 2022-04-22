package model

type Blacklist struct {
	ID           int64  `gorm:"column:id;type:BIGINT;autoIncrement;comment:'唯一id';NOT NULL" json:"id"`
	UserName     string `gorm:"column:username;type:VARCHAR(128);index;comment:'用户名'" json:"username"`
	Enable       *bool  `gorm:"column:enable;comment:'是否启用'" json:"enable"`
	Mode         int8   `gorm:"column:mode;comment:'模式 1不记录 2记录'" json:"mode"`
	MatchPattern int8   `gorm:"column:match_pattern;comment:'模式 1域名通配符 2正则'" json:"match_pattern"`
	Rules        string `gorm:"column:rules;comment:'匹配规则一行一个'" json:"rules"`
	Base
}
