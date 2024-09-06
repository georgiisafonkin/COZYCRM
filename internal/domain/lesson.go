package domain

import "time"

type Lesson struct {
	Id int64
	Content *LessonContent
	InsertTimestamp time.Time
	UpdateTimestamp time.Time
}

type LessonContent struct {
	Topic string
	Body string // остерегаться большого содержимого
}
