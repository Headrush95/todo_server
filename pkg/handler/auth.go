package handler

import (
	"github.com/gin-gonic/gin"
	first_server "github.com/headrush95/first_server"
	"net/http"
)

// @Summary SignUp
// @Tags Auth
// @Description Create new account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body first_server.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input first_server.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags Auth
// @Description Login
// @ID login
// @Accept json
// @Produce json
// @Param input body signInInput true "credentials"
// @Success 200 {integer} string "token"
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
