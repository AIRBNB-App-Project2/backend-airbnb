package auth

import "be/entities"

type Auth interface {
	Login(UserLogin entities.User) (entities.User, error)
}
