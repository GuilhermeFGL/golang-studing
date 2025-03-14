package repository

import (
	"database/sql"
	"example.com/m/v2/API/src/models"
	"log"
)

type User struct {
	db *sql.DB
}

// NewRepository create a new repository for user
func NewRepository(db *sql.DB) *User {
	return &User{db}
}

// CreateUser create and return a new user
func (repository User) CreateUser(user models.User) (models.User, error) {
	statement, err := repository.db.Prepare("INSERT INTO users (name, email, nickname, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return models.User{}, err
	}

	exec, err := statement.Exec(user.Name, user.Email, user.NickName, user.Password)
	if err != nil {
		return models.User{}, err
	}

	user.ID, err = exec.LastInsertId()

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Fatal("Error closing statement")
		}
	}(statement)

	return user, nil
}

// FetchUser find user by id
func (repository User) FetchUser(id uint64) (models.User, error) {
	var rows, err = repository.db.Query("SELECT id, name, nickname, email, created_at FROM users WHERE id = ?", id)
	if err != nil {
		return models.User{}, err
	}

	if rows.Next() {
		var user models.User

		if err = rows.Scan(&user.ID, &user.Name, &user.NickName, &user.Email, &user.CreatedAt); err == nil {
			return user, nil
		} else {
			return models.User{}, err
		}
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("Error closing query")
		}
	}(rows)

	return models.User{}, nil
}

// SearchUser return all user by filter
func (repository User) SearchUser(name string) ([]models.User, error) {
	var rows *sql.Rows
	var err error
	if name != "" {
		rows, err = repository.db.Query("SELECT id, name, nickname, email, created_at FROM users WHERE lower(name) LIKE ? or lower(nickname) LIKE ?", "%"+name+"%", "%"+name+"%")
	} else {
		rows, err = repository.db.Query("SELECT id, name, nickname, email, created_at FROM users")
	}

	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(&user.ID, &user.Name, &user.NickName, &user.Email, &user.CreatedAt); err == nil {
			users = append(users, user)
		} else {
			return nil, err
		}
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("Error closing query")
		}
	}(rows)

	return users, nil
}

// UpdateUser update and return user by id
func (repository User) UpdateUser(id uint64, updateUser models.User) (models.User, error) {
	foundUser, err := repository.FetchUser(id)
	if err != nil || foundUser.ID == 0 {
		return models.User{}, err
	}

	statement, err := repository.db.Prepare("update users set name = ?, nickname = ?, email = ? where id = ?")
	if err != nil {
		return models.User{}, err
	}

	if _, err = statement.Exec(updateUser.Name, updateUser.NickName, updateUser.Email, id); err != nil {
		return models.User{}, err
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Fatal("Error closing statement")
		}
	}(statement)

	foundUser.Name = updateUser.Name
	foundUser.NickName = updateUser.NickName
	foundUser.Email = updateUser.Email
	return foundUser, nil
}

// DeleteUser delete user by id
func (repository User) DeleteUser(id uint64) (bool, error) {
	statement, err := repository.db.Prepare("delete from users where id = ?")
	if err != nil {
		return false, err
	}

	if rowsAffected, err := statement.Exec(id); err == nil {
		if rowAffected, _ := rowsAffected.RowsAffected(); rowAffected > 0 {
			return true, nil
		} else {
			return false, nil
		}
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Fatal("Error closing query")
		}
	}(statement)

	return false, nil
}
