package graph

import (
	"context"
	"yells-post/graph/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, allowComments bool) (*model.Post, error) {
    newPost := &model.Post{
        ID:            "generated-id", // Можно сгенерировать ID, например, через github.com/google/uuid
        Title:         title,
        Content:       content,
        AllowComments: allowComments,
        Comments:      []*model.Comment{},
    }
    // Здесь можно добавить сохранение в базу или in-memory хранилище.
    return newPost, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, postID string, parentID *string, text string) (*model.Comment, error) {
    newComment := &model.Comment{
        ID:       "generated-comment-id",
        Text:     text,
        Author:   "anonymous", // или извлекать данные об авторе из контекста
        ParentID: parentID,
        Replies:  []*model.Comment{},
    }
    // Здесь можно добавить логику сохранения комментария и привязки к посту.
    return newComment, nil
}
