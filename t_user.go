package main

import (
	"database/sql"
	"time"
)

const FIND_TUSER_BY_ID = "SELECT * FROM t_user WHERE user_id = $1"
const CREATE_TUSER = "INSERT INTO t_user VALUES ($1, $2, $3, $4, $5, $6)"

// Структура физических пользователей Телеграма
type TUser struct {
	UserID       int64
	UserName     string
	FirstName    string
	LastName     string
	Lang         string
	CreationDate time.Time
}

func getTUserByID(user_id int64) (TUser, error) {
	var tu TUser
	err := db.QueryRow(FIND_TUSER_BY_ID, user_id).Scan(&tu.UserID, &tu.UserName, &tu.FirstName, &tu.LastName, &tu.Lang, &tu.CreationDate)
	if err == sql.ErrNoRows {
		return tu, nil
	} else if err != nil {
		return tu, err
	}
	return tu, nil
}

func createTUser(tu TUser) error {
	_, err := db.Exec(CREATE_TUSER, tu.UserID, tu.UserName, tu.FirstName, tu.LastName, tu.Lang, tu.CreationDate)
	return err
}
