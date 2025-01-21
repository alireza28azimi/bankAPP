package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"main.go/entity"
)

func (d *MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {

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
func (d *MysqlDB) Register(u entity.User) (entity.User, error) {
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
func (d *MysqlDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	row := d.db.QueryRow(`select *from user where phone_number =?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, fmt.Errorf("you phone number is not register")
		}
		return entity.User{}, fmt.Errorf("unexpected error %w", err)
	}
	return user, nil

}
func (d *MysqlDB) GetUserByID(userID uint) (entity.User, error) {
	row := d.db.QueryRow(`SELECT id, phone_number, name, password, created_at FROM users WHERE id = ?`, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, err
		}
		return entity.User{}, err

	}
	return user, nil
}
