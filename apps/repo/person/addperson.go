package person

import "github.com/wulungtry/techtalk/entity"

type AddRepo interface {
	Add(*entity.Person) error
}
