package handler

import (
	"github.com/gin-gonic/gin"
	First_server "github.com/headrush95/first_server"
	"net/http"
	"strconv"
)

// @Summary Create new item
// @Security ApiKeyAuth
// @Tags Items
// @Description Create new item in list
// @ID create-item
// @Accept json
// @Produce json
// @Param input body First_server.TodoItem true "item info"
// @Param id path int true "List ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input First_server.TodoItem

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get all items in list
// @Security ApiKeyAuth
// @Tags Items
// @Description Get all items in list
// @ID get-all-items
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Success 200 {object} []First_server.TodoItem
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/lists/{id}/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}

// @Summary Get item by ID
// @Security ApiKeyAuth
// @Tags Items
// @Description Get item by ID
// @ID get-item-by-id
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} First_server.TodoItem
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

// @Summary Update item
// @Security ApiKeyAuth
// @Tags Items
// @Description Update item's info
// @ID update-item
// @Accept json
// @Produce json
// @Param input body First_server.UpdateItemInput true "updated item info"
// @Param id path int true "Item ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input First_server.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete item
// @Security ApiKeyAuth
// @Tags Items
// @Description Delete item
// @ID delete-item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
