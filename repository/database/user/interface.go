package user

import "be/entities"

type User interface {
	Create(user entities.User) (entities.User, error)
	GetById(user_uid string) (GetByIdResponse, error)
	Update(user_uid string, upUser entities.User) (entities.User, error)
	Delete(userUid string) (entities.User, error)
}
