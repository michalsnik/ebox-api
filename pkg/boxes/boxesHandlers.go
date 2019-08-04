package boxes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BoxesHandlers interface {
	GetBoxById (c *gin.Context)
	PutBox (c *gin.Context)
}

type boxesHandlers struct {
	BoxesHandlers
	svc BoxesService
}

type CreateBoxRequest struct {
	Name string `json:"name"`
}

func NewHandlers(svc BoxesService) BoxesHandlers {
	return &boxesHandlers{svc: svc}
}

func (h *boxesHandlers) GetBoxById (c *gin.Context) {
	boxId, err := strconv.Atoi(c.Param("boxID"))

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	box := h.svc.GetBoxById(boxId)

	c.JSON(http.StatusOK, box)
}

func (h *boxesHandlers) PutBox (c *gin.Context) {
	var payload CreateBoxRequest

	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	box, err := h.svc.CreateBox(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, box)
}