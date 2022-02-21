package user

import "be/entities"

type User interface {
	Create(user entities.User) (entities.User, error)
}
