package models

import (
	"errors"

	"github.com/aungsannphyo/go-restapi/db"
	"github.com/aungsannphyo/go-restapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Signup() error {
	query := `INSERT INTO 
	users(email, password)
	VALUES (?, ?)`

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	user.ID = userId
	return err
}

func (user *User) Login() error {
	query := "SELECT id,password FROM users WHERE email = ? "
	row := db.DBInstance.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsBalid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsBalid {
		return errors.New("Credentials invalid")
	}

	return nil
}
