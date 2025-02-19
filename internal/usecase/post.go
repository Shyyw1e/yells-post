package usecase

import (
	"errors"
	"log/slog"

	"yells-post/internal/domain"
	"yells-post/graph/model"
)

type PostUsecase struct {
	Repo domain.PostRepository
}

func NewPostUsecase(repo domain.PostRepository) *PostUsecase {
	return &PostUsecase{
		Repo: repo,
	}
}

func (u *PostUsecase) CreatePost(title, content string, allowComments bool) (*model.Post, error) {
	if title == "" || content == ""{
		err := errors.New("заголовок и содержание не должны быть пустыми")
		slog.Error("Ошибка создания поста", err)
		return nil, err
	}

	post := &model.Post{
		Title: title,
		Content: content,
		AllowComments: allowComments,
	}

	createdPost, err := u.Repo.CreatePost(post)
	if err != nil {
		slog.Error("Ошибка создания поста в репозитории", err, "title", title)
		return nil, err
	}

	slog.Info("Пост успешно создан", "id", createdPost.ID)
	return createdPost, nil
}

func (u *PostUsecase) GetPost(id string) (*model.Post, error) {
	post, err := u.Repo.GetPost(id)
	if err != nil {
		slog.Error("Ошибка получения поста", err, "id", id)
		return nil, err
	}
	return post, nil
}

func (u *PostUsecase) ListPosts(page, pageSize int) ([]*model.Post, error) {
	posts, err := u.Repo.ListPosts(page, pageSize)
	if err != nil {
		slog.Error("Ошибка получения списка постов", err, "page", page, "pageSize", pageSize)
		return nil, err
	}

	return posts, nil
}

func (u *PostUsecase) UpdatePost(post *model.Post) error {
	err := u.Repo.UpdatePost(post)
	if err != nil {
		slog.Error("Ошибка обновления поста", err, "id", post.ID)
	}
	return nil
}