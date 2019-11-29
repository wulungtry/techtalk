package controller

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/wulungtry/techtalk/api/proto/person"
	"github.com/wulungtry/techtalk/api/proto/response"
	"github.com/wulungtry/techtalk/apps/usecase"
	"github.com/wulungtry/techtalk/common"
)

type personalServiceServer struct {
	db *gorm.DB
}

func NewPersonalServiceServer(db *gorm.DB) person.PersonalServiceServer {
	return &personalServiceServer{
		db: db,
	}
}

func (p *personalServiceServer) AddPerson(ctx context.Context, model *person.AddPersonModel) (*response.ResponseBase, error) {
	addCase := usecase.NewPersonCase(p.db)

	var result = response.ResponseBase{}

	if err := addCase.Add(model); err != nil {
		result.Status = response.ResponseBase_FAIL
		result.Message = common.InsertFail("person")
		result.Data = nil

		return &result, err
	}
	result.Status = response.ResponseBase_SUCCESS
	result.Message = common.InsertSuccess("person")
	result.Data = nil

	return &result, nil
}
