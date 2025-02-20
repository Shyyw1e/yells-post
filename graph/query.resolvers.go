package graph

import (
	"context"
	"yells-post/graph/model"
)

func (r *queryResolver) Posts(ctx context.Context, page *int32, pageSize *int32) ([]*model.Post, error) {
    var p, ps int
    if page != nil {
        p = int(*page)
    } else {
        p = 1
    }
    if pageSize != nil {
        ps = int(*pageSize)
    } else {
        ps = 10
    }
    return r.PostUsecase.ListPosts(p, ps)
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
    return r.PostUsecase.GetPost(id)
}

func (r *Resolver) Query() QueryResolver {
    return &queryResolver{r}
}

type queryResolver struct { *Resolver}
