package person

import "github.com/wulungtry/techtalk/entity"

type UpdateRepo interface {
	Update(person *entity.Person) error
}
