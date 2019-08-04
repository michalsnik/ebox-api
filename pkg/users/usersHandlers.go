package users

import (
	"ebox-api/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersHandlers interface {
	PostUser (c *gin.Context)
	GetMe (c *gin.Context)
}

type usersHandlers struct {
	svc UsersService
}

type PostUserRequestData struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	AvatarURL string `json:"avatarURL"`
}

func NewUsersHandlers(svc UsersService) UsersHandlers {
	return &usersHandlers{
		svc: svc,
	}
}

func (h *usersHandlers) PostUser (c *gin.Context) {
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

func (h *usersHandlers) GetMe (c *gin.Context) {
	userID := c.Value("userID").(int)

	user, err := h.svc.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, response.Create(nil, err))
		return
	}

	c.JSON(http.StatusOK, response.Create(user, nil))
}