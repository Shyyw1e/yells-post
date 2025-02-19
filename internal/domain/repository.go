package domain

import (
	"yells-post/graph/model"
)

type PostRepository interface {
	CreatePost(post *model.Post) (*model.Post, error)

	GetPost(id string) (*model.Post, error)

	ListPosts(page, pageSize int) ([]*model.Post, error)

	UpdatePost(post *model.Post) error
}

type CommentRepository interface {
	CreateComment(postID string, comment *model.Comment) (*model.Comment, error)

	ListCommentsbyPost(postID string, page, pageSize int) ([]*model.Comment, error)
}