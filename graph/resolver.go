package graph

import "yells-post/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	PostUsecase *usecase.PostUsecase
	CommentUsecase *usecase.CommentUsecase
}
