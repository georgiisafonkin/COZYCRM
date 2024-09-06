package domain

ManyToMany таблица ОБУЧЕНИЯ
user_id <-> course_id
date_joined
rating

type User struct {
	юзернейм
	дата регистрации
}

User.getCourses()
