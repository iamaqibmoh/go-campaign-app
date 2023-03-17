package repository

import "bwa-campaign-app/model/domain"

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
}
