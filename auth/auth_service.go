package auth

import (
	"histos-backend/database"

	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string, dbinstance database.DatabaseConnection) (string, error) {
	row := dbinstance.Connection.QueryRow("SELECT userid, passhash from Users WHERE email=$1", email)
	var passhash string
	var userId string
	err := row.Scan(&userId, &passhash)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passhash), []byte(password))
	if err != nil {
		return "", err
	}

	jwt, err := GenerateJwt(userId)

	if err != nil {
		return "", err
	}

	return jwt, nil
}

func Signup(email string, password string, dbinstance database.DatabaseConnection) error {
	passhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = dbinstance.Connection.Exec("INSERT INTO Users values(DEFAULT, $1, $2)", email, passhash)

	if err != nil {
		return err
	}

	return nil
}
