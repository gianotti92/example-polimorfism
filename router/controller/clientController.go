package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type InteractWay interface {
	Interact() string
}

type ClientController struct {
	ClientType string
	Interact   InteractWay
}

func (clientController ClientController) Client(c *gin.Context) {
	var value = clientController.Interact.Interact()

	c.String(http.StatusCreated, value)
}
