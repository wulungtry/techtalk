package repo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
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

func (pr *PersonRepo) Update(p *entity.Person) error {
	if err := pr.db.Table("person").Save(p).Error; err != nil {
		return err
	}

	return nil
}

func (pr *PersonRepo) GetAll() []domain.Person {
	var result []domain.Person
	pr.db.Table("person").Scan(&result)

	return result
}

func (pr *PersonRepo) GetById(id string) domain.Person {
	var result domain.Person
	pr.db.Table("person").Where("personalid = ?", id).Limit(1).Scan(&result)
	//pr.db.Table("person").Select().Where(&entity.Person{Personalid: id}).FirstOrInit(&result)

	return result
}
