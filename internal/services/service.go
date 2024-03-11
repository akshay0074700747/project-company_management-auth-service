package services

import (
	"context"

	"github.com/akshay0074700747/project-company_management-auth-service/entities"
	"github.com/akshay0074700747/project-company_management-auth-service/helpers"
	"github.com/akshay0074700747/project-company_management-auth-service/internal/usecases"
	"github.com/akshay0074700747/projectandCompany_management_protofiles/pb/authpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthServiceServer struct {
	Usecase usecases.AuthUsecaseInterfaces
	authpb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(usecase usecases.AuthUsecaseInterfaces) *AuthServiceServer {
	return &AuthServiceServer{
		Usecase: usecase,
	}
}

func (auth *AuthServiceServer) LoginUser(ctx context.Context, req *authpb.LoginUserRequest) (*authpb.LoginUserRes, error) {

	res, err := auth.Usecase.LoginUser(entities.Authentication{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		helpers.PrintErr(err, "error occured at loggin user usecase...")
		return nil, err
	}

	return &authpb.LoginUserRes{
		UserID:  res.UserID,
		IsAdmin: res.IsAdmin,
	}, nil

}

func (auth *AuthServiceServer) InsertUser(ctx context.Context, req *authpb.InsertUserReq) (*emptypb.Empty, error) {

	_, err := auth.Usecase.InsertUser(entities.Authentication{
		UserID:   req.UserID,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		helpers.PrintErr(err, "error at insert user usecase")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
