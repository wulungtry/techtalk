package person

import "github.com/wulungtry/techtalk/api/proto/person"

type AddCase interface {
	Add(model *person.AddPersonModel) error
}
