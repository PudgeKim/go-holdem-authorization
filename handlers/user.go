package handlers

import (
	"github.com/Pudgekim/application"
	"github.com/Pudgekim/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetUserById(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("id")

	interactor := application.UserInteractor{Repository: h.userRepo}

	user, err := interactor.GetUser(ctx, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      user.Id,
		"name":    user.Name,
		"email":   user.Email,
		"balance": user.Balance,
	})

}

func (h *Handler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var req entity.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	interactor := application.UserInteractor{Repository: h.userRepo}

	newUser := entity.NewUser(req.Id, req.Name, req.Email, 0)
	if err := interactor.AddUser(ctx, newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      newUser.Id,
		"name":    newUser.Name,
		"email":   newUser.Email,
		"balance": newUser.Balance,
	})
}
