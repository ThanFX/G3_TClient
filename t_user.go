package main

import (
	"database/sql"
	"time"
)

const dbFindTUserByID = "SELECT * FROM t_user WHERE user_id = $1"
const dbCreateTUser = "INSERT INTO t_user VALUES ($1, $2, $3, $4, $5, $6)"

// TUser - Структура физических пользователей Телеграма
type TUser struct {
	UserID       int64
	UserName     string
	FirstName    string
	LastName     string
	Lang         string
	CreationDate time.Time
}

func getTUserByID(userID int64) (TUser, error) {
	var tu TUser
	err := db.QueryRow(dbFindTUserByID, userID).Scan(&tu.UserID, &tu.UserName, &tu.FirstName, &tu.LastName, &tu.Lang, &tu.CreationDate)
	if err == sql.ErrNoRows {
		return tu, nil
	} else if err != nil {
		return tu, err
	}
	return tu, nil
}

func createTUser(tu TUser) error {
	_, err := db.Exec(dbCreateTUser, tu.UserID, tu.UserName, tu.FirstName, tu.LastName, tu.Lang, tu.CreationDate)
	return err
}
