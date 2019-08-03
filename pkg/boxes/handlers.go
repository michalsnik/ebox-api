package boxes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handlers struct {
	svc *Service
}

type CreateBoxRequest struct {
	Name string `json:"name"`
}

func NewHandlers(svc *Service) *Handlers {
	return &Handlers{svc: svc}
}

func (h *Handlers) GetBoxById (c *gin.Context) {
	boxId, err := strconv.Atoi(c.Param("boxID"))

	if err != nil {
		return
	}

	box := h.svc.GetBoxById(boxId)

	c.JSON(http.StatusOK, box)
}

func (h *Handlers) PutBox (c *gin.Context) {
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