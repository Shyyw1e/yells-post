package postgres

import (
	"errors"

	"log/slog"

	"github.com/jmoiron/sqlx"
	"yells-post/internal/domain"
	"yells-post/graph/model"
)

type Repo struct {
	DB *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) CreatePost(post *model.Post) (*model.Post, error) {
	query := `
		INSERT INTO posts (title, content, allow_comments)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	var id string
	err := r.DB.QueryRow(query, post.Title, post.Content, post.AllowComments).Scan(&id)
	if err != nil {
		slog.Error("Ошибка создания поста", err, "title", post.Title)
		return nil, err
	}
	post.ID = id
	slog.Info("Пост успешно создан", "id", post.ID)
	return post, nil
}

func (r *Repo) GetPost(id string) (*model.Post, error) {
	query := `
		SELECT id, title, content, allow_comments
		FROM posts
		WHERE id = $1;
	`

	var post model.Post
	err := r.DB.Get(&post, query, id)
	if err != nil {
		slog.Error("Ошибка получения поста", err, "id", id)
		return nil, err
	}
	return &post, nil
}

func (r *Repo) ListPosts(page, pageSize int) ([]*model.Post, error) {
	offset := (page - 1) * pageSize
	query := `
		SELECT id, title, content, allow_comments
		FROM posts
		ORDER BY id
		LIMIT $1 OFFSET $2;
	`

	var posts []*model.Post
	err := r.DB.Select(&posts, query, pageSize, offset)
	if err != nil {
		slog.Error("Ошибка получения списка постов", err, "page", page, "pageSize", pageSize)
		return nil, err
	}
	return posts, nil
}

func (r *Repo) UpdatePost(post *model.Post) error {
	query := `
		UPDATE posts
		SET title = $1, content = $2, allow_comment = $3
		WHERE id = $4;
	`

	result, err := r.DB.Exec(query, post.Title, post.Content, post.AllowComments, post.ID)
	if err != nil {
		slog.Error("Ошибка обновления поста", err, "id", post.ID)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		slog.Error("Ошибка проверки количества обновленных строк", err, "id", post.ID)
		return err
	}
	if rows == 0 {
		err = errors.New("пост не найден")
		slog.Error("Ошибка", err, "id", post.ID)
		return err
	}
	slog.Info("Пост успешно обновлен", "id", post.ID)
	return nil
}

func (r *Repo) CreateComment(postID string, comment *model.Comment) (*model.Comment, error) {
	query := `
		INSERT INTO comments (post_id, text, author, parent_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	var id string
	err := r.DB.QueryRow(query, postID, comment.Text, comment.Author, comment.ParentID).Scan(&id)
	if err != nil {
		slog.Error("Ошибка создания комментария", err, "postID", postID)
		return nil, err
	}
	return comment, nil
}

func (r *Repo) ListCommentsbyPost(postID string, page, pageSize int) ([]*model.Comment, error) {
	offset := (page - 1) * pageSize
	query := `
		SELECT id, text, author, parent_id
		FROM comments
		WHERE post_id = $1
		ORDER BY id
		LIMIT $2 OFFSET $3;
	`

	var comments []*model.Comment
	err := r.DB.Select(&comments, query, postID, pageSize, offset)
	if err != nil {
		slog.Error("Ошибка получения комментариев к посту", err, "postID", postID)
		return nil, err
	}
	return comments, nil
}

var _ domain.PostRepository = (*Repo)(nil)
var _ domain.CommentRepository = (*Repo)(nil)