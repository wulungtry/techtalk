package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/wulungtry/techtalk/apps/domain"
	"github.com/wulungtry/techtalk/entity"
)

type PersonRepo struct {
	db *gorm.DB
}

func NewPersonRepo(db *gorm.DB) *PersonRepo {
	return &PersonRepo{
		db: db,
	}
}

func (pr *PersonRepo) Add(p *entity.Person) error {
	if err := pr.db.Create(p).Error; err != nil {
		return err
	}

	return nil
}

func (pr *PersonRepo) Update(person *entity.Person) error {
	if err := pr.db.Update(person).Error; err != nil {
		return err
	}

	return nil
}

func (pr *PersonRepo) GetAll() []domain.Person {
	var result []domain.Person
	pr.db.Table("person").Find(&result)

	return result
}

func (pr *PersonRepo) GetById(id string) domain.Person {
	var result domain.Person
	pr.db.Table("person").First(&result, id)

	return result
}
