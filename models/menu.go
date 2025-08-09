package models

import (
	"context"
	"time"

	"github.com/mratri10/go-rich/config"
)

type Menu struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	AdminId int       `json:"adminId"`
	MenuId  int       `json:"menuId"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
type RoleMenu struct {
	ID      int       `json:"id"`
	RoleId  int       `json:"roleId"`
	AdminId int       `json:"adminId"`
	MenuId  int       `json:"menuId"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func CreateMenu(name string, adminId int, parent int) error {
	_, error := config.DB.Exec(context.Background(),
		"INSERT INTO menus (name, admin_id, menu_id) VALUES ($1,$2, $3)",
		name, adminId, parent)
	return error
}

func GetMenu() ([]Menu, error) {
	rows, err := config.DB.Query(context.Background(),
		"SELECT id, name, admin_id, menu_id from menus")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []Menu
	for rows.Next() {
		var m Menu
		if err := rows.Scan(&m.ID, &m.Name, &m.AdminId, &m.MenuId); err != nil {
			return nil, err
		}
		menus = append(menus, m)
	}
	return menus, err
}
func GetMenuById(id int) (Menu, error) {
	var m Menu
	err := config.DB.QueryRow(context.Background(),
		`SELECT id, name, admin_id, menu_id from menus where id=$1`,
		id).Scan(&m.ID, &m.Name, &m.AdminId, &m.MenuId)
	return m, err
}
func UpdateMenu(id int, name string, adminId int, menuId int) (Menu, error) {
	_, err := config.DB.Exec(context.Background(),
		`UPDATE menus SET name=$1, updated_at=NOW(), admin_id=$2, menu_id=$3
		where id=$4`, name, adminId, menuId, id)
	if err != nil {
		return Menu{}, err
	}
	return GetMenuById(id)
}
func DeleteMenu(id int) error {
	_, err := config.DB.Exec(context.Background(),
		`DELETE FROM menus where id=$1`, id)
	return err
}

func CreateMenuRole(MenuId int, UserId int, AdminId int) error {
	_, error := config.DB.Exec(context.Background(),
		"INSERT INTO rolemenu (role_id, admin_id, menu_id) VALUES ($1, $2, $3)",
		UserId, AdminId, MenuId)
	return error
}

func DeleteMenuRole(id int) error {
	_, err := config.DB.Exec(context.Background(),
		`DELETE FROM usermenu where id=$1`, id)
	return err
}
