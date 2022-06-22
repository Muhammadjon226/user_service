package v1

import (
	"net/http"

	"github.com/Muhammadjon226/user_service/models"
	"github.com/Muhammadjon226/user_service/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

//ListUsers method for get list of users
// @Summary List Users
// @Description ListUsers API is for get list of users
// @Tags user
// @Accept  json
// @Produce  json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} models.ListUserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/list-users/ [get]
func (h *HandlerV1) ListUsers(c *gin.Context) {

	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)

	if errStr != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	response, err := h.userService.ListUsers(context.Background(), &models.ListUserRequest{
		Limit: params.Limit,
		Page:  params.Page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response)

}
