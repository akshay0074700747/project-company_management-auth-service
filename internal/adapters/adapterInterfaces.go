package adapters

import "github.com/akshay0074700747/project-company_management-auth-service/entities"

type UserAdapterInterfaces interface {
	LoginUser(entities.Authentication) (entities.Authentication, error)
	InsertUser(entities.Authentication) (entities.Authentication, error)
	// InsertUserintoAuthorization(entities.Authorization) (entities.Authorization, error)
	AuthorizeUser(user_id string) (bool, error)
	VerifyPass(string,string) (bool,error)
}
