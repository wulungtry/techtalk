package person

import "github.com/wulungtry/techtalk/apps/domain"

type GetRepo interface {
	GetAll() []domain.Person
	GetById(id string) domain.Person
}
