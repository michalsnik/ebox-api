package auth

import (
	"ebox-api/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandlers struct {
	svc *AuthService
}

type SignInRequestData struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type SignInResponseData struct {
	Token string `json:"token"`
}

func NewAuthHandlers(svc *AuthService) *AuthHandlers {
	return &AuthHandlers{
		svc: svc,
	}
}

func (h *AuthHandlers) SignIn (c *gin.Context) {
	var reqData SignInRequestData

	err := c.BindJSON(&reqData)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Create(nil, err))
		return
	}

	token, err := h.svc.AuthenticateUser(reqData.Email, reqData.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Create(nil, err))
		return
	}

	resData := &SignInResponseData{
		Token: token,
	}

	c.JSON(http.StatusOK, response.Create(resData, nil))
}

