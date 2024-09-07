package domain

ManyToMany таблица ОБУЧЕНИЯ
user_id <-> course_id
date_joined
rating

type User struct {
	username String
	signUpDate String
}

User.getCourses()
