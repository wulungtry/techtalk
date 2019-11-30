package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/wulungtry/techtalk/api/proto/person"
	"github.com/wulungtry/techtalk/apps/usecase"
	"github.com/wulungtry/techtalk/common"
)

type personalServiceServer struct {
	db *gorm.DB
}

func NewPersonalServiceServer(db *gorm.DB) personal.PersonalServiceServer {
	return &personalServiceServer{
		db: db,
	}
}

func (p *personalServiceServer) GetPersonById(ctx context.Context, request *personal.GetByIdModel) (*personal.GetResponseModel, error) {
	getCase := usecase.NewPersonCase(p.db)
	var result = personal.GetResponseModel{}

	dto := getCase.GetById(request.PersonalId)

	result.Name = dto.Name
	result.CityOfBirth = dto.CityOfBirth
	result.DateOfBirth, _ = ptypes.TimestampProto(dto.DateOfBirth)
	result.Weight = dto.Weight
	result.Height = dto.Height
	result.IsMarried = dto.IsMarried

	return &result, nil
}

func (p *personalServiceServer) GetAllPerson(context.Context, *personal.GetAllModel) (*personal.GetAllResponseModel, error) {
	getCase := usecase.NewPersonCase(p.db)
	var result = personal.GetAllResponseModel{}

	dto := getCase.GetAll()

	for _, i := range dto {
		var p personal.GetResponseModel
		p.Name = i.Name
		p.CityOfBirth = i.CityOfBirth
		p.DateOfBirth, _ = ptypes.TimestampProto(i.DateOfBirth)
		p.Height = i.Height
		p.Weight = i.Weight
		p.IsMarried = i.IsMarried

		result.Data = append(result.Data, &p)
	}

	return &result, nil
}

func (p *personalServiceServer) AddPerson(ctx context.Context, model *personal.AddPersonModel) (*personal.ResponseModel, error) {
	addCase := usecase.NewPersonCase(p.db)
	var result = personal.ResponseModel{}

	if err := addCase.Add(model); err != nil {
		result.Status = personal.ResponseModel_FAIL
		result.Message = common.InsertFail("person")
		result.Data = nil

		return &result, err
	}
	result.Status = personal.ResponseModel_SUCCESS
	result.Message = common.InsertSuccess("person")
	result.Data = nil

	return &result, nil
}

func (p *personalServiceServer) UpdatePerson(ctx context.Context, model *personal.UpdatePersonModel) (*personal.ResponseModel, error) {
	updCase := usecase.NewPersonCase(p.db)
	var result = personal.ResponseModel{}

	if err := updCase.Update(model); err != nil {
		result.Status = personal.ResponseModel_FAIL
		result.Message = common.UpdateFail("person")
		result.Data = nil

		return &result, err
	}

	result.Status = personal.ResponseModel_SUCCESS
	result.Message = common.UpdateSuccess("person")
	result.Data = nil

	return &result, nil
}
