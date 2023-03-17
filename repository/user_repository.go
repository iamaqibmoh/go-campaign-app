package repository

import "bwa-campaign-app/model/domain"

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	FindByID(id int) (domain.User, error)
	Update(user domain.User) (domain.User, error)
}
