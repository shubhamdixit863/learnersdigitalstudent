package repository

import (
	"assesment_5/internal/models"
	"errors"
)

type MemoryRepo struct {
	posts  []models.Post
	nextID int
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		posts:  []models.Post{},
		nextID: 1,
	}
}

func (m *MemoryRepo) GetAll() []models.Post {
	return m.posts
}

func (m *MemoryRepo) Create(post models.Post) models.Post {
	post.ID = m.nextID
	m.nextID++
	m.posts = append(m.posts, post)
	return post
}

func (m *MemoryRepo) Update(id int, updated models.Post) (models.Post, error) {
	for i, p := range m.posts {
		if p.ID == id {
			updated.ID = id
			m.posts[i] = updated
			return updated, nil
		}
	}
	return models.Post{}, errors.New("post not found")
}

func (m *MemoryRepo) Delete(id int) error {
	for i, p := range m.posts {
		if p.ID == id {
			m.posts = append(m.posts[:i], m.posts[i+1:]...)
			return nil
		}
	}
	return errors.New("post not found")
}
