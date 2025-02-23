package usecase_test

import (
	"errors"
	"testing"

	"yells-post/graph/model"
	"yells-post/internal/domain"
	"yells-post/internal/usecase"
)

type dummyPostRepo struct {
	posts map[string]*model.Post
}

func newDummyPostRepo() *dummyPostRepo {
	return &dummyPostRepo{posts: make(map[string]*model.Post)}
}

func (r *dummyPostRepo) CreatePost(post *model.Post) (*model.Post, error) {
	post.ID = "test-id"
	r.posts[post.ID] = post
	return post, nil
}

func (r *dummyPostRepo) GetPost(id string) (*model.Post, error) {
	post, ok := r.posts[id]
	if !ok {
		return nil, errors.New("post not found")
	}
	return post, nil
}

func (r *dummyPostRepo) ListPosts(page, pageSize int) ([]*model.Post, error) {
	var list []*model.Post
	for _, p := range r.posts {
		list = append(list, p)
	}
	return list, nil
}

func (r *dummyPostRepo) UpdatePost(post *model.Post) error {
	if _, ok := r.posts[post.ID]; !ok {
		return errors.New("post not found")
	}
	r.posts[post.ID] = post
	return nil
}

var _ domain.PostRepository = (*dummyPostRepo)(nil)

func TestCreatePost(t *testing.T) {
	repo := newDummyPostRepo()
	usecasePost := usecase.NewPostUsecase(repo)

	post, err := usecasePost.CreatePost("Test Title", "Test Content", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if post.ID != "test-id" {
		t.Errorf("unexpected post id 'test-id', got %v", post.ID)
	}
}

func TestCreatePost_EmptyFields(t *testing.T) {
	repo := newDummyPostRepo()
	usecasePost := usecase.NewPostUsecase(repo)

	_, err := usecasePost.CreatePost("", "Content", true)
	if err == nil {
		t.Error("expected error for empty title , got nil")
	}
}

func TestCreatePost_EmptyContent(t *testing.T) {
	repo := newDummyPostRepo()
	usecasePost := usecase.NewPostUsecase(repo)

	_, err := usecasePost.CreatePost("Title", "", true)
	if err == nil {
		t.Error("expected error for empty content , got nil")
	}
}