package model

type Word struct {
	Word string `gorm:"column:word;primaryKey;type:VARCHAR(128);uniqueIndex;comment:'词'" json:"word"`
}
