package inmemory

import (
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"yells-post/graph/model"
	"yells-post/internal/domain"
)

type InMemoryRepo struct {
	mu sync.RWMutex //читателей может быть много, а автор 1
	posts map[string]*model.Post // [id]пост
	comments map[string][]*model.Comment // [id]слайс_комментов
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		posts: make(map[string]*model.Post),
		comments: make(map[string][]*model.Comment),
	}
}

func (r *InMemoryRepo) CreatePost(post *model.Post) (*model.Post, error){ //создание поста
	r.mu.Lock()
	defer r.mu.Unlock()

	post.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	if post.Comments == nil {
		post.Comments = []*model.Comment{}
	}

	r.posts[post.ID] = post
	slog.Info("Пост создан", "id", post.ID)
	return post, nil
}

func (r *InMemoryRepo) GetPost(id string) (*model.Post, error) { //получение по id
	r.mu.RLock()
	defer r.mu.RUnlock()

	post, exists := r.posts[id]
	if !exists {
		err := errors.New("пост не найден")
		slog.Error("Ошибка получения поста", err, "id", id)
		return nil, err
	}

	return post, nil
}

func (r *InMemoryRepo) ListPosts(page, pageSize int) ([]*model.Post, error) { //пагинация
	r.mu.RLock()
	defer r.mu.RUnlock()

	var posts []*model.Post

	for _, p := range r.posts {
		posts = append(posts, p)
	}

	offset := (page - 1) * pageSize 
	if offset >= len(posts) {
		return []*model.Post{}, nil
	}
	end := offset + pageSize
	if end > len(posts){
		end = len(posts)
	}
	return posts[offset:end], nil
}

func (r *InMemoryRepo) UpdatePost(post *model.Post) error { //обновление поста
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.posts[post.ID]; !exists {
		err := errors.New("пост не найден")
		slog.Error("Ошибка обновления поста", err, "id", post.ID)
		return err
	}

	r.posts[post.ID] = post
	slog.Info("Пост обновлен", "id", post.ID)
	return nil
}

//теперь комменты

func (r *InMemoryRepo) CreateComment(postID string, comment *model.Comment) (*model.Comment, error){ //создание поста
	r.mu.Lock()
	defer r.mu.Unlock()

	comment.ID = fmt.Sprintf("%d", time.Now().UnixNano())


	post, exists := r.posts[postID]
	if !exists {
		err := errors.New("пост для комментариев не найден")
		slog.Error("Ошибка создания комментария", err, "id", postID)
	}
	if !post.AllowComments {
		err:= errors.New("комментарии запрещены для этого поста")
		slog.Error("Ошибка создания комментария", err, "postID", post.ID)
		return nil, err
	}

	r.comments[post.ID] = append(r.comments[post.ID], comment)
	post.Comments = append(post.Comments, comment)
	slog.Info("Комментарий создан", "id", comment.ID, "postID", post.ID)
	return comment, nil
}

func (r *InMemoryRepo) ListCommentsbyPost(postID string, page, pageSize int) ([]*model.Comment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	comments, exists := r.comments[postID]
	if !exists {
		return []*model.Comment{}, nil
	}
	offset := (page - 1) * pageSize
	if offset >= len(comments) {
		return []*model.Comment{}, nil
	}

	end := offset + pageSize
	if end > len(comments) {
		end = len(comments)
	}
	
	return comments[offset:end], nil
}


var _ domain.PostRepository = (*InMemoryRepo)(nil)
var _ domain.CommentRepository = (*InMemoryRepo)(nil)

