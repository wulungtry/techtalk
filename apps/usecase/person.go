package usecase

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/wulungtry/techtalk/api/proto/person"
	"github.com/wulungtry/techtalk/apps/repo"
	"github.com/wulungtry/techtalk/entity"
)

type PersonCase struct {
	db *gorm.DB
}

func NewPersonCase(db *gorm.DB) *PersonCase {
	return &PersonCase{
		db: db,
	}
}

func (pc *PersonCase) Add(model *person.AddPersonModel) error {
	var entity = entity.Person{}

	entity.Personalid = model.PersonalId
	entity.Name = model.Name
	entity.Cityofbirth = model.CityOfBirth
	entity.Dateofbirth, _ = ptypes.Timestamp(model.DateOfBirth)
	entity.Height = model.Height
	entity.Weight = model.Weight
	entity.Ismarried = model.IsMarried

	repo := repo.NewPersonRepo(pc.db)
	if err := repo.Add(&entity); err != nil {
		return err
	}

	return nil
}
