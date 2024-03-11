package usecases

import (
	"errors"

	"github.com/akshay0074700747/project-company_management-auth-service/entities"
	"github.com/akshay0074700747/project-company_management-auth-service/helpers"
	"github.com/akshay0074700747/project-company_management-auth-service/internal/adapters"
)

type AuthUseCases struct { 
	Adapter adapters.UserAdapterInterfaces
}

func NewAuthUseCases(adapter adapters.UserAdapterInterfaces) *AuthUseCases {
	return &AuthUseCases{
		Adapter: adapter,
	}
}

func (auth *AuthUseCases) InsertUser(req entities.Authentication) (entities.Authentication, error) {

	var err error
	req.Password, err = helpers.Hash_pass(req.Password)
	if err != nil {
		helpers.PrintErr(err, "error occured at hashing password...")
		return entities.Authentication{}, err
	}

	res, err := auth.Adapter.InsertUser(req)
	if err != nil {
		helpers.PrintErr(err, "error occured at isert user adapter...")
		return entities.Authentication{}, err
	}

	return res, nil
}

func (auth *AuthUseCases) LoginUser(req entities.Authentication) (entities.Authorization, error) {

	ress, err := auth.Adapter.LoginUser(req)
	if err != nil {
		helpers.PrintErr(err, "error at logging in the user")
		return entities.Authorization{}, err
	}
	isAdmin, err := auth.Adapter.AuthorizeUser(ress.UserID)
	if err != nil {
		helpers.PrintErr(err, "error occured at checking isdmin user adapter...")
		return entities.Authorization{}, err
	}
	if isAdmin {
		verified, err := auth.Adapter.VerifyPass(ress.Email, req.Password)
		if err != nil {
			helpers.PrintErr(err, "errror occued at VerifyPass adapter")
			return entities.Authorization{}, err
		}
		if !verified {
			return entities.Authorization{}, errors.New("please enter a valid email,password")
		}
	} else {
		if err := helpers.VerifyPassword(ress.Password, req.Password); err != nil {
			helpers.PrintErr(err, "eror occured at VerifyPassword")
			return entities.Authorization{}, errors.New("please enter a valid email,password")
		}
	}

	return entities.Authorization{
		UserID:  ress.UserID,
		IsAdmin: isAdmin,
	}, nil
}
