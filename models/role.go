package models

import (
	"context"
	"time"

	"github.com/mratri10/go-rich/config"
)

type Role struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	AdminId int       `json:"adminId"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
type UserRole struct {
	ID      int       `json:"id"`
	UserId  int       `json:"user_id"`
	AdminId int       `json:"admin_id"`
	RoleId  int       `json:"role_id"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func CreatedRole(name string, adminId int) error {
	_, error := config.DB.Exec(context.Background(),
		"INSERT INTO roles (name, admin_id) VALUES ($1, $2)",
		name, adminId)
	return error
}
func GetRole() ([]Role, error) {
	rows, err := config.DB.Query(context.Background(),
		"SELECT id,name,admin_id from roles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []Role
	for rows.Next() {
		var r Role
		if err := rows.Scan(&r.ID, &r.Name, &r.AdminId); err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	return roles, err
}
func GetRoleById(id int) (Role, error) {
	var r Role
	err := config.DB.QueryRow(context.Background(),
		"SELECT id, name, admin_id FROM roles where id=$1",
		id).Scan(&r.ID, &r.Name, &r.AdminId)
	return r, err

}
func UpdateRole(id int, name string, adminId int) (Role, error) {

	_, err := config.DB.Exec(context.Background(),
		`UPDATE roles SET name=$1, updated_at=NOW(), admin_id=$2 
		where id=$3`,
		name, adminId, id)

	if err != nil {
		return Role{}, err
	}
	return GetRoleById(id)
}
func DeleteRole(id int) error {
	_, err := config.DB.Exec(context.Background(),
		"DELETE FROM roles where id=$1", id)
	return err
}
func CreateRoleUser(RoleId int, UserId int, AdminId int) error {
	_, err := config.DB.Exec(context.Background(),
		"INSERT INTO userrole (user_id, admin_id, role_id) VALUES ($1,$2,$3)",
		UserId, AdminId, RoleId)
	return err
}

func DeleteRoleUser(id int) error {
	_, err := config.DB.Exec(context.Background(),
		`DELETE FROM userrole where id=$id`, id)
	return err
}
