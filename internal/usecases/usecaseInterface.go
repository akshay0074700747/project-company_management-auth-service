package usecases

import "github.com/akshay0074700747/project-company_management-auth-service/entities"

type AuthUsecaseInterfaces interface {
	InsertUser(entities.Authentication) (entities.Authentication, error)
	LoginUser(entities.Authentication) (entities.Authorization, error)
}
