package usecase

import (
	"errors"
	"log/slog"

	"yells-post/internal/domain"
	"yells-post/graph/model"
)

type CommentUsecase struct {
	Repo domain.CommentRepository
}

func NewCommentUsecase(repo domain.CommentRepository) *CommentUsecase {
	return &CommentUsecase{
		Repo: repo,
	}
}

func (u *CommentUsecase) CreateComment(postID, text string, parentID *string) (*model.Comment, error) {
	if text == "" {
		err := errors.New("комментарий не может быть пустым")
		slog.Error("Ошибка создания комментария: пустой текст", err)
		return nil, err
	}

	if len(text) > 2000 {
		err := errors.New("длина комментария не может превышать 2000 символов")
		slog.Error("Ошибка создания комментария: превышена длина", err)
		return nil, err
	}

	comment := &model.Comment{
		Text: text,
		Author: "anonymous",
		ParentID: parentID,
	}

	createdComment, err := u.Repo.CreateComment(postID, comment)
	if err != nil {
		slog.Error("Ошибка создания комментария в репозитории", err, "postID", postID)
		return nil, err
	}

	slog.Info("Комментарий успешно создан", "id", createdComment.ID, "postID", postID)
	return createdComment, nil
}

func (u *CommentUsecase) ListCommentsbyPost(postID string, page, pageSize int) ([]*model.Comment, error) {
	comments, err := u.Repo.ListCommentsbyPost(postID, page, pageSize)
	if err != nil {
		slog.Error("Ошибка получения комментариев для поста", err, "postID", postID)
		return nil, err
	}
	return comments, nil
}