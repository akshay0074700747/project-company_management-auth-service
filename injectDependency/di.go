package injectdependency

import (
	"github.com/akshay0074700747/project-company_management-auth-service/config"
	"github.com/akshay0074700747/project-company_management-auth-service/db"
	"github.com/akshay0074700747/project-company_management-auth-service/internal/adapters"
	"github.com/akshay0074700747/project-company_management-auth-service/internal/services"
	"github.com/akshay0074700747/project-company_management-auth-service/internal/usecases"
)

func Initialize(cfg config.Config) *services.AuthEngine {

	db := db.ConnectDB(cfg)
	adapter := adapters.NewAuthAdapter(db)
	usecase := usecases.NewAuthUseCases(adapter)
	server := services.NewAuthServiceServer(usecase)

	go adapters.Adminise(db, cfg)

	return services.NewAuthEngine(server)
}
