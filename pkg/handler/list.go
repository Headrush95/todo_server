package handler

import (
	"github.com/gin-gonic/gin"
	First_server "github.com/headrush95/first_server"
	"net/http"
	"strconv"
)

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags Lists
// @Description Create new todo list
// @ID create-list
// @Accept json
// @Produce json
// @Param input body First_server.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input First_server.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	// записываем id списка в тело ответа
	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []First_server.TodoList `json:"data"`
}

// @Summary Get all user's todo lists
// @Security ApiKeyAuth
// @Tags Lists
// @Description Get todo lists
// @ID get-all-lists
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary Get todo list by ID
// @Security ApiKeyAuth
// @Tags Lists
// @Description Get todo list by ID
// @ID get-list-by-id
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Success 200 {integer} integer First_server.TodoList
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
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

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Update todo list
// @Security ApiKeyAuth
// @Tags Lists
// @Description Update todo list's info
// @ID update-list
// @Accept json
// @Produce json
// @Param input body First_server.UpdateListInput true "updated list info"
// @Param id path int true "List ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
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
	var input First_server.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete todo list
// @Security ApiKeyAuth
// @Tags Lists
// @Description Delete todo list
// @ID delete-list
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
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

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
