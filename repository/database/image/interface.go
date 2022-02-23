package image

import "be/entities"

type Image interface {
	Create(image entities.Image) (entities.Image, error)
}
