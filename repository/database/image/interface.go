package image

import "be/entities"

type Image interface {
	Create(image string) (entities.Image, error)
}
