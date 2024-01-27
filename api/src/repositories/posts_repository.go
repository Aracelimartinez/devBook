package repositories

import (
	"api/src/models"
	"database/sql"
)

// Representa um repositório de publicações
type posts struct {
	db *sql.DB
}

// Cria um repositorio de publicações
func NewPostRepository(db *sql.DB) *posts {
	return &posts{db}
}

// Insere uma publicação no banco de dados
func (p posts) CreatePost(newPost *models.Post) (uint64, error) {
	stmt, err := p.db.Prepare(
		"INSERT INTO posts (title, content, author_id) VALUES($1, $2, $3) RETURNING id",
	)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	var postID uint64
	err = stmt.QueryRow(newPost.Title, newPost.Content, newPost.AuthorID).Scan(&postID)
	if err != nil {
		return 0, err
	}

	return postID, nil
}

// Traz publicações do banco de dados dos usuários seguidos e também do próprio usuário que fez a requisição
func (p posts) SearchPosts(userID uint64) (*[]models.Post, error) {
	rows, err := p.db.Query(
		`SELECT DISTINCT p.*, u.nick
		FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		INNER JOIN followers f ON f.user_id = p.author_id
		WHERE u.id = $1
		OR f.follower_id = $2
		ORDER BY p.created_at DESC`, userID, userID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNickname,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return &posts, nil
}

// Traz uma única publicação do banco de dados
func (p posts) SearchByID(postID uint64) (*models.Post, error) {
	var post models.Post

	stmt, err := p.db.Prepare(
		`SELECT p.*, u.nick
		FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		WHERE p.id = $1  `,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(postID)
	err = row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.AuthorID,
		&post.Likes,
		&post.CreatedAt,
		&post.AuthorNickname,
	)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// Altera as informações de um post no banco de dados
func (u posts) UpdatePost(postID uint64, updatedPost *models.Post) error {
	stmt, err := u.db.Prepare(
		"UPDATE posts SET title = $1, content = $2 WHERE id = $3",
	)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(updatedPost.Title, updatedPost.Content, postID); err != nil {
		return err
	}

	return nil
}

// Exclui uma publicação do banco de dados
func (p posts) DeletePost(postID uint64) error {
	stmt, err := p.db.Prepare("DELETE FROM posts WHERE id = $1")
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}

// Traz todas as publicações no banco de dados de um usuário específico
func (p posts) SearchByUser(userID uint64) (*[]models.Post, error) {
	rows, err := p.db.Query(
		`SELECT p.*, u.nick
		FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		WHERE p.author_id = $1
		ORDER BY p.created_at DESC`, userID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNickname,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return &posts, nil
}

// Adiciona una curtida a um post no banco de dados
func (u posts) LikePost(postID uint64) error {
	stmt, err := u.db.Prepare(
		"UPDATE posts SET likes = likes + 1 WHERE id = $1",
	)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}

// Subtrai una curtida a um post no banco de dados
func (u posts) UnLikePost(postID uint64) error {
	stmt, err := u.db.Prepare(
		`UPDATE posts SET likes =
		CASE
		WHEN likes > 0 THEN likes - 1
		ELSE 0
		END
		WHERE id = $1`,
	)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}
