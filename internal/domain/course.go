package domain

обучения в user.go, здесь учеников прошедших курс и обучающихся не хранить

type Course struct {
	courseId int64
	title String
	category String
	lessons String[]
	finishedCount int64
	inProgressCount int64
	rating flout
	creationDate String
}

Course.getAuthors() []User
