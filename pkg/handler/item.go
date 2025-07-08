package handler

import (
	"net/http"
	"strconv"

	"github.com/akhmadeevrus/todo-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllItems(c *gin.Context) {
}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id poram")
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err = h.services.TodoItem.Create(userId, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateItem(c *gin.Context) {
}

func (h *Handler) deleteItem(c *gin.Context) {
}
