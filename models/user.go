package models

import (
	"fmt"

	"github.com/wanderporto/api-json-crud/utils"
)

type User struct {
	Id        int        `json:"id"`
	Name      NullString `json:"name"`
	Email     NullString `json:"email"`
	Password  NullString `json:"password"`
	CreatedAt NullString `json:"createdAt"`
}

func NewUser(user User) (bool, error) {
	con := Connect()

	hash, err := utils.Hash(user.Password.String)

	sql := `insert into users(name,email,password) values(?,?,?)`

	stmt, err := con.Prepare(sql)

	if err != nil {
		fmt.Printf(err.Error())
		return false, err
	}

	user.Password.String = fmt.Sprintf("%s", hash)

	_, err = stmt.Exec(&user.Name.String, &user.Email.String, &user.Password.String)

	if err != nil {
		fmt.Printf(err.Error())
		return false, err
	}

	defer con.Close()
	defer stmt.Close()

	return true, nil

}

func GetUsers() ([]User, error) {

	con := Connect()

	sql := `select id,name,email,password,createdat from users`

	rs, err := con.Query(sql)

	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}

	var users []User

	for rs.Next() {

		var user User

		err := rs.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

		if err != nil {
			fmt.Printf(err.Error())
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUser(id int) (User, error) {

	var user User
	con := Connect()

	sql := `select id,name,email,password,createdat from users where id = ?`

	rs := con.QueryRow(sql, id)

	err := rs.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return user, err
	}

	defer con.Close()

	return user, nil
}

func UpdateUser(user User) (int64, error) {
	con := Connect()

	sql := `update users
			   set name = ?,
				   email = ?,
				   password = ?
			 where id = ?`

	stmt, err := con.Prepare(sql)

	if err != nil {
		fmt.Printf("error Prepare User: %s", err)
		return 0, err
	}

	rs, err := stmt.Exec(user.Name.String, user.Email.String, user.Password.String, user.Id)

	if err != nil {
		fmt.Printf("error scan: %s", err)
		return 0, err
	}

	rowsAffected, err := rs.RowsAffected()

	defer con.Close()
	defer stmt.Close()

	return rowsAffected, nil

}

func DeleteUser(id int) (int64, error) {
	con := Connect()

	sql := `delete from users where id = ?`

	stmt, err := con.Prepare(sql)

	if err != nil {
		return 0, err
	}

	rs, err := stmt.Exec(id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := rs.RowsAffected()

	if err != nil {
		return 0, err
	}

	defer con.Close()
	defer stmt.Close()

	return rowsAffected, nil
}

func GetUserByEmail(email string) (User, error) {
	con := Connect()

	sql := `select id,name,email,password,createdat from users where email = ?`

	rs, err := con.Query(sql, email)

	if err != nil {
		return User{}, err
	}

	var user User
	if rs.Next() {
		err := rs.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

		if err != nil {
			return User{}, err
		}
	}

	defer con.Close()
	defer rs.Close()

	return user, nil

}

func Signin(email, password string) (User, error) {

	user, err := GetUserByEmail(email)

	if err != nil {
		return User{}, err
	}

	err = utils.VerifyPassword([]byte(user.Password.String), []byte(password))

	if err != nil {
		return User{}, err
	}

	return user, nil
}
