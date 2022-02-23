package user

import "be/entities"

type User interface {
	Create(user entities.User) (entities.User, error)
	GetById(useruid string) (entities.User, error)
	Update(userUid string, newUser entities.User) (entities.User, error)
	Delete(userUid string) error
}
