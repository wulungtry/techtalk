package usecase

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	personal "github.com/wulungtry/techtalk/api/proto/person"
	"github.com/wulungtry/techtalk/apps/domain"
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

func (pc *PersonCase) Add(model *personal.AddPersonModel) error {
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

func (pc *PersonCase) Update(model *personal.UpdatePersonModel) error {
	var entity = entity.Person{}

	entity.Personalid = model.PersonalId
	entity.Name = model.Name
	entity.Cityofbirth = model.CityOfBirth
	entity.Dateofbirth, _ = ptypes.Timestamp(model.DateOfBirth)
	entity.Height = model.Height
	entity.Weight = model.Weight
	entity.Ismarried = model.IsMarried

	repo := repo.NewPersonRepo(pc.db)
	if err := repo.Update(&entity); err != nil {
		return err
	}

	return nil
}

func (pc *PersonCase) GetAll() []domain.Person {
	var result = []domain.Person{}
	repo := repo.NewPersonRepo(pc.db)

	result = repo.GetAll()

	return result
}

func (pc *PersonCase) GetById(id string) domain.Person {
	var result = domain.Person{}
	repo := repo.NewPersonRepo(pc.db)

	result = repo.GetById(id)

	return result
}
