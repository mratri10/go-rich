package models

import (
	"context"
	"time"

	"github.com/mratri10/go-rich/config"
)

type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	AdminId  int       `json:"adminId"`
	Name     string    `json:"name"`
	Password string    `json:"-"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func CreatedUser(username, hashPassword, name string, adminId int) error {
	_, err := config.DB.Exec(context.Background(),
		"INSERT INTO users (username, password, name, admin_id) VALUES ($1, $2, $3, $4)",
		username, hashPassword, name, adminId)
	return err
}

func GetUserByUsername(username string) (User, error) {
	var u User
	err := config.DB.QueryRow(context.Background(),
		"SELECT id,name, username FROM users WHERE username=$1",
		username).Scan(&u.ID, &u.Name, &u.Username)
	return u, err
}
func GetUserAll() ([]User, error) {
	rows, err := config.DB.Query(context.Background(),
		"SELECT id, name, username, admin_id, created_at, updated_at from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Username, &u.AdminId, &u.Updated, &u.Created); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, err
}

func UpdateUserByUsername(username, name, password string, userId int) (User, error) {
	_, err := config.DB.Exec(context.Background(),
		"UPDATE users SET password=$1, name=$2, updated_at=NOW(), admin_id=$4 where username = $3",
		password, name, username, userId)
	if err != nil {
		return User{}, err
	}
	return GetUserByUsername(username)
}

func DeleteUseById(id int) error {
	_, err := config.DB.Exec(context.Background(),
		"DELETE FROM users where id=$1", id)
	return err
}
