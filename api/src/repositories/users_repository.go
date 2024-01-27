package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Representa um repositório de usuários
type users struct {
	db *sql.DB
}

// Cria um repositorio de usuários
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Insere um usuário no banco de dados
func (u users) CreateUser(newUser *models.User) (uint64, error) {
	stmt, err := u.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	var userID uint64
	err = stmt.QueryRow(newUser.Name, newUser.Nick, newUser.Email, newUser.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// Traz todos os usuários que atendem um filtro de nome ou nickname
func (u users) SearchUsers(nameOrNick string) (*[]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := u.db.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE name ILIKE $1 OR nick ILIKE $2", nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}

// Traz um usuário do banco de dados
func (u users) SearchUserByID(userID uint64) (*models.User, error) {
	var user models.User

	stmt, err := u.db.Prepare(
		"SELECT id, name, nick, email, created_at FROM users WHERE id = $1  ",
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(userID)
	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Nick,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Altera as informações de um usuário no banco de dados
func (u users) UpdateUser(userID uint64, updatedUser *models.User) error {
	stmt, err := u.db.Prepare(
		"UPDATE users SET name = $1, nick = $2, email = $3 WHERE id = $4",
	)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(updatedUser.Name, updatedUser.Nick, updatedUser.Email, userID); err != nil {
		return err
	}

	return nil
}

// Exclui as informações de um usuário no banco de dados
func (u users) DeleteUser(userID uint64) error {
	stmt, err := u.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(userID); err != nil {
		return err
	}

	return nil
}

// Busca um usuário por email e retorna o seu ID e senha com hash
func (u users) SearchUserByEmail(email string) (*models.User, error) {
	var user models.User

	stmt, err := u.db.Prepare(
		"SELECT id, password FROM users WHERE email = $1",
	)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(email)
	err = row.Scan(&user.ID, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Permite que um usuario siga outro
func (u users) Follow(userID, followerID uint64) error {
	stmt, err := u.db.Prepare(
		`INSERT INTO followers (user_id, follower_id)
		VALUES ($1, $2)
		ON CONFLICT (user_id, follower_id) DO NOTHING`,
	)
	if err != nil {
		return err
	}

	defer stmt.Close()
	if _, err = stmt.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Permite que um usuariopare de seguir outro
func (u users) Unfollow(userID, followerID uint64) error {
	stmt, err := u.db.Prepare("DELETE FROM followers WHERE user_id = $1 AND follower_id = $2")
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Traz todos os seguidores de um usuário
func (u users) SearchFollowers(userID uint64) (*[]models.User, error) {
	rows, err := u.db.Query(
		`SELECT id, u.name, u.nick, u.email, u.created_at
		FROM users u
		INNER JOIN followers f ON u.id = f.follower_id
		WHERE f.user_id = $1`, userID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var followers []models.User

	for rows.Next() {
		var follower models.User
		err = rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return &followers, nil
}

// Traz todos usuários que um determinado usuário está seguindo
func (u users) SearchUsersFollowed(userID uint64) (*[]models.User, error) {
	rows, err := u.db.Query(
		`SELECT id, u.name, u.nick, u.email, u.created_at
		FROM users u
		INNER JOIN followers f ON u.id = f.user_id
		WHERE f.follower_id = $1`, userID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var usersFollowed []models.User

	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		usersFollowed = append(usersFollowed, user)
	}

	return &usersFollowed, nil
}

// Busca a senha de um usuário pelo ID
func (u users) SearchPassword(userID uint64) (string, error) {
	var user models.User

	stmt, err := u.db.Prepare(
		"SELECT password FROM users WHERE id = $1",
	)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	row := stmt.QueryRow(userID)
	err = row.Scan(
		&user.Password,
	)
	if err != nil {
		return "", err
	}

	return user.Password, nil
}

// Altera a senha de um usuário
func (u users) UpdatePassword(userID uint64, updatedPassword string) error {
	stmt, err := u.db.Prepare(
		"UPDATE users SET password = $1 WHERE id = $2",
	)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(updatedPassword, userID); err != nil {
		return err
	}

	return nil
}
