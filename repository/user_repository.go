package repository

import (
	"github.com/Kundhavi2798/grpc-user-service/db"
	"github.com/Kundhavi2798/grpc-user-service/models"
)

func CreateUser(user models.User) (int64, error) {
	res, err := db.DB.Exec(
		"INSERT INTO users(name,email) VALUES(?,?)",
		user.Name,
		user.Email,
	)
	if err != nil {
		return 0, nil
	}
	return res.LastInsertId()
}

func GetUser(id int32) (*models.User, error) {
	row := db.DB.QueryRow("SELECT id,name,email FROM users WHERE id= ?", id)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
