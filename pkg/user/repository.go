package user

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	database *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{database: db}
}

func (r *userRepository) RegisterUser(username string, hashedPw string) (err error) {
	var InsertUser string = "INSERT INTO users (username, password) VALUES (?, ?)"

	res, err := r.database.Exec(InsertUser, username, hashedPw)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return errors.New("falied to register user")
	}

	return
}

func (r *userRepository) LogInUser(username string, password string) (userId int, err error) {
	var LogInUser string = "SELECT id, password FROM users WHERE username = ?"

	var hashedPw string
	if err = r.database.QueryRow(LogInUser, username).Scan(&userId, &hashedPw); err != nil {
		return 0, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(password)); err != nil {
		return 0, err
	}

	return int(userId), nil
}

func (r *userRepository) RetrieveCredit(userId int) (currentCredit int, err error) {
	var SelectCredit string = "SELECT credit FROM users WHERE id = ?"

	if err = r.database.QueryRow(SelectCredit, userId).Scan(&currentCredit); err != nil {
		return 0, err
	}

	return currentCredit, nil
}

func (r *userRepository) InsertCredit(userId int, credit int64) (newCredit int, err error) {
	var UpdateCredit string = "UPDATE users SET credit = credit + ? WHERE id = ?"

	res, err := r.database.Exec(UpdateCredit, credit, userId)
	if err != nil {
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	} else if affected == 0 {
		return 0, errors.New("no affected rows")
	}

	var RetrieveCredit string = "SELECT credit FROM users WHERE id = ?"
	if err = r.database.QueryRow(RetrieveCredit, userId).Scan(&newCredit); err != nil {
		return 0, err
	}

	return
}
