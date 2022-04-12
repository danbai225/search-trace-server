package modle

import "time"

type Trace struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Url        string    `json:"url"`
	Content    string    `json:"content"`
	BrowseTime time.Time `json:"browse_time"`
}
