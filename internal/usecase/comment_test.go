package usecase_test

import (
	"errors"
	"testing"

	"yells-post/graph/model"
	"yells-post/internal/domain"
	"yells-post/internal/usecase"
)

type dummyCommentRepo struct {
	comments map[string][]*model.Comment
	posts map[string]*model.Post
}

func newDummyCommentRepo() *dummyCommentRepo {
	return &dummyCommentRepo{
		comments: make(map[string][]*model.Comment),
		posts: map[string]*model.Post{},
	}
}

func (r *dummyCommentRepo) CreateComment(postID string, comment *model.Comment) (*model.Comment, error) {
	if _, ok := r.posts[postID]; !ok {
		return nil, errors.New("post not found")
	}
	comment.ID = "comment-test-id"
	r.comments[postID] = append(r.comments[postID], comment)
	return comment, nil
}

func (r *dummyCommentRepo) ListCommentsbyPost(postID string, page, pageSize int) ([]*model.Comment, error) {
	return r.comments[postID], nil
}

var _ domain.CommentRepository = (*dummyCommentRepo)(nil)

func TestCreateComment(t *testing.T) {
	repo := newDummyCommentRepo()
	repo.posts["1"] = &model.Post{
		ID: "1",
		Title: "test Post",
		Content: "content",
		AllowComments: true,
	}

	usecaseComent := usecase.NewCommentUsecase(repo)

	comment, err := usecaseComent.CreateComment("1", "Test comment", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if comment.ID != "comment-test-id" {
		t.Errorf("expected comment id 'comment-test-id', got %v", comment.ID)
	}
}

func TestCreateComment_EmptyContent(t *testing.T) {
	repo := newDummyCommentRepo()
	repo.posts["1"] = &model.Post{
		ID: "1",
		Title: "test Post",
		Content: "content",
		AllowComments: true,
	}
	usecaseComment := usecase.NewCommentUsecase(repo)
	_, err := usecaseComment.CreateComment("1", "", nil)
	if err == nil {
		t.Errorf("expected error for empty comment text, got nil")
	}
}

