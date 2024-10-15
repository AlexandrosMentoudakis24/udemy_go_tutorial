package models

import (
	"backend/db"
	"backend/utils"
	"errors"
)

type User struct {
	ID		 int64
	Email	 string `binding:"required`
	Password string `binding:"required`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result , err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}
	
	userId, err := result.LastInsertId()

	u.ID = userId

	if err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateCredentials() error {
	query := `
	SELECT id, password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Invalid credentials!")
	}

	isValidPassword := utils.ComparePasswords(u.Password, retrievedPassword)

	if !isValidPassword {
		return errors.New("Invalid credentials!")
	}

	return nil
}

