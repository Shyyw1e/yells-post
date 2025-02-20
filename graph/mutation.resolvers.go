package graph

import (
	"context"
	"yells-post/graph/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, allowComments bool) (*model.Post, error) {
    return r.PostUsecase.CreatePost(title, content, allowComments)
}

func (r *mutationResolver) CreateComment(ctx context.Context, postID string, parentID *string, text string) (*model.Comment, error) {
    return r.CommentUsecase.CreateComment(postID, text, parentID)
}

func (r *Resolver) Mutation() MutationResolver {
    return &mutationResolver{r}
}

type mutationResolver struct { *Resolver}
