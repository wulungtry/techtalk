package person

import personal "github.com/wulungtry/techtalk/api/proto/person"

type AddCase interface {
	Add(model *personal.AddPersonModel) error
}
