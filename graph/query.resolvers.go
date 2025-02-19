package graph

import (
	"context"
	"yells-post/graph/model"
)

func (r *queryResolver) Posts(ctx context.Context, page *int32, pageSize *int32) ([]*model.Post, error) {
    // Пример с тестовыми данными. Можно добавить обработку пагинации.
    posts := []*model.Post{
        {
            ID:            "1",
            Title:         "Первый пост",
            Content:       "Содержимое первого поста",
            AllowComments: true,
            Comments:      []*model.Comment{}, // пока пустой список
        },
        {
            ID:            "2",
            Title:         "Второй пост",
            Content:       "Содержимое второго поста",
            AllowComments: false,
            Comments:      []*model.Comment{},
        },
    }
    return posts, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
    // Пример с тестовыми данными.
    return &model.Post{
        ID:            id,
        Title:         "Тестовый пост",
        Content:       "Содержимое тестового поста",
        AllowComments: true,
        Comments:      []*model.Comment{},
    }, nil
}
