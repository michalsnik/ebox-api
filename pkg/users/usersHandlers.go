package users

import (
	"ebox-api/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersHandlers struct {
	svc *UsersService
}

type PostUserRequestData struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	AvatarURL string `json:"avatarURL"`
}

func NewUsersHandlers(svc *UsersService) *UsersHandlers {
	return &UsersHandlers{
		svc: svc,
	}
}

func (h *UsersHandlers) PostUser (c *gin.Context) {
	var reqData PostUserRequestData

	err := c.BindJSON(&reqData)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Create(nil, err))
		return
	}

	user, err := h.svc.CreateUser(reqData)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Create(nil, err))
		return
	}

	c.JSON(http.StatusCreated, response.Create(user, nil))
}
