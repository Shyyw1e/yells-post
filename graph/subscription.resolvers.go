package graph

import (
    "context"
    "time"
    "yells-post/graph/model"
)

func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *model.Comment, error) {
    commentChan := make(chan *model.Comment, 1)

    go func() {
        select {
        case <-ctx.Done():
            close(commentChan)
            return
        case <-time.After(5 * time.Second):
            commentChan <- &model.Comment{
                ID:       "subscribed-comment-id",
                Text:     "Новый комментарий через подписку",
                Author:   "subscriber",
                ParentID: nil,
                Replies:  []*model.Comment{},
            }
            close(commentChan)
        }
    }()

    return commentChan, nil
}