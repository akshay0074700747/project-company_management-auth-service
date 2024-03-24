package adapters

import (
	"errors"

	"github.com/akshay0074700747/project-company_management-auth-service/config"
	"github.com/akshay0074700747/project-company_management-auth-service/entities"
	"github.com/akshay0074700747/project-company_management-auth-service/helpers"
	"gorm.io/gorm"
)

type AuthAdapter struct {
	DB *gorm.DB
}

func NewAuthAdapter(db *gorm.DB) *AuthAdapter {
	return &AuthAdapter{
		DB: db,
	}
}

func (auth *AuthAdapter) LoginUser(req entities.Authentication) (entities.Authentication, error) {

	query := "SELECT * FROM authentications WHERE email = $1"
	var res entities.Authentication
	result := auth.DB.Raw(query, req.Email).Scan(&res)

	if result.Error != nil {
		return res, result.Error
	}

	if result.RowsAffected == 0 {
		return res, errors.New("no user found with the given credentials")
	}
	return res, nil
}

func (auth *AuthAdapter) InsertUser(req entities.Authentication) (entities.Authentication, error) {

	query := "INSERT INTO authentications (user_id,email,password) VALUES($1,$2,$3) RETURNING user_id,email"
	var res entities.Authentication

	tx := auth.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := auth.DB.Raw(query, req.UserID, req.Email, req.Password).Scan(&res).Error; err != nil {
		tx.Rollback()
		return res, err
	}

	if err := tx.Commit().Error; err != nil {
		return res, err
	}

	return res, nil
}

// func (auth *AuthAdapter) InsertUserintoAuthorization(req entities.Authorization) (entities.Authorization, error) {

// 	query := "INSERT INTO authorizations (user_id,is_admin) VALUES($1,$2) RETURNING user_id,is_admin"
// 	var res entities.Authorization

// 	tx := auth.DB.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	if err := auth.DB.Raw(query, req.UserID, req.IsAdmin).Scan(&res).Error; err != nil {
// 		tx.Rollback()
// 		return res, err
// 	}

// 	if err := tx.Commit().Error; err != nil {
// 		return res, err
// 	}

// 	return res, nil
// }

func (auth *AuthAdapter) AuthorizeUser(user_id string) (bool, error) {

	query := "SELECT is_admin FROM authorizations WHERE user_id = $1"
	var res bool

	if err := auth.DB.Raw(query, user_id).Scan(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (auth *AuthAdapter) VerifyPass(email, password string) (bool, error) {

	query := "SELECT password FROM authentications WHERE email = $1 AND password = $2"
	var pass string

	if err := auth.DB.Raw(query, email, password).Scan(&pass).Error; err != nil {
		return false, err
	}

	if pass != password {
		return false, errors.New("the passwords doesnt match")
	}

	return true, nil
}

func Adminise(db *gorm.DB, conf config.Config) {

	query := "SELECT * FROM authorizations WHERE is_admin = true"

	tx := db.Exec(query)
	if tx.Error != nil {
		helpers.PrintErr(tx.Error, "error at checking admin...")
		return
	}

	if tx.RowsAffected == 0 {
		id := helpers.GenUuid()
		query = "INSERT INTO authentications (user_id,email,password) VALUES($1,$2,$3)"
		if err := db.Exec(query, id, conf.AdminEmail, conf.AdminPass).Error; err != nil {
			helpers.PrintErr(err, "error happened at creating admin")
		}

		query = "INSERT INTO authorizations (user_id,is_admin) VALUES($1,$2)"
		if err := db.Exec(query, id, true).Error; err != nil {
			helpers.PrintErr(err, "error happened at creating admin")
		}
	}
}
