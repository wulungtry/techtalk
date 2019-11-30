package person

import "github.com/wulungtry/techtalk/apps/domain"

type GetCase interface {
	GetAll() []domain.Person
	GetById(id string) domain.Person
}
