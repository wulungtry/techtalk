package person

import personal "github.com/wulungtry/techtalk/api/proto/person"

type UpdCase interface {
	Update(model *personal.UpdatePersonModel) error
}
