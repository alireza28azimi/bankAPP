package mysql

import (
	"database/sql"
	"fmt"
	"main.go/entity"
)

func (d *MySqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {

	row := d.db.QueryRow(`select *from users where phone_number =?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}
	return false, nil

}
func (d *MySqlDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`INSERT INTO users(name, phone_number, password) VALUES(?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)

	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command %w", err)
	}
	//error is always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil
}
func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt []uint8
	var user entity.User

	err := row.Scan(&user.ID, &user.PhoneNumber, &user.Name, &user.Password, &createdAt)
	if err != nil {
		return entity.User{}, fmt.Errorf("record not found ")
	}
	return user, nil

}
