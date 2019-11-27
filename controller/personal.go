package controller

import (
	"context"
	"github.com/wulungtry/techtalk/api/proto/person"
	"github.com/wulungtry/techtalk/api/proto/response"
)

type personalServiceServer struct{}

func NewPersonalServiceServer() person.PersonalServiceServer {
	return &personalServiceServer{}
}

func (p personalServiceServer) AddPerson(ctx context.Context, model *person.AddPersonModel) (*response.ResponseBase, error) {
	return &response.ResponseBase{
		Status:  response.ResponseBase_SUCCESS,
		Message: "Berhasil",
		Data:    nil,
	}, nil
}
