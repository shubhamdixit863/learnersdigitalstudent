package handlers

import (
	"assesment_5/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, h.repo.GetAll())
}

func (h *Handler) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil || post.Title == "" || post.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post input"})
		return
	}

	created := h.repo.Create(post)

	token, err := h.jwtService.CreateJWT(strconv.Itoa(created.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"post":  created,
		"token": token,
	})
}

func (h *Handler) UpdatePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.repo.Update(id, post)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
