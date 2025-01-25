package controller

import (
	"kassech/backend/pkg/constants"
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueueManagerRouteController struct {
	Service *service.QueueManagerRouteService
}

func NewQueueManagerRouteController(Service *service.QueueManagerRouteService) *QueueManagerRouteController {
	return &QueueManagerRouteController{Service: Service}
}

func (c *QueueManagerRouteController) CreateRoute(ctx *gin.Context) {
	var req service.CreateQueueManagerRouteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	route, err := c.Service.CreateRoute(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, route)
}
func (c *QueueManagerRouteController) GetAllRoutes(ctx *gin.Context) {

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}

	perPage, err := strconv.Atoi(ctx.Query("per_page"))
	if err != nil {
		perPage = 10
	}

	search := ctx.Query("search")

	roles, _ := ctx.Get("role")

	// Filter routes by user ID if the role is not Admin and userID is passed
	if !utils.Contains(roles.([]string), constants.QMRoleName) && ctx.Query("user_id") != "" {
		userIDUint, err := strconv.ParseUint(ctx.Query("user_id"), 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		routes, total, err := c.Service.GetAllRoutesForUser(uint(userIDUint), page, perPage, search)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch routes"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":  routes,
			"total": total,
		})
		return
	}

	routes, total, err := c.Service.GetAllRoutes(page, perPage, search)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch routes"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  routes,
		"total": total,
	})
}

func (c *QueueManagerRouteController) GetRoute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	_, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Implement similar pattern to GetAll but with single ID
	// You'll need to add a GetByID method in repository/service
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

func (c *QueueManagerRouteController) DeleteRoute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.Service.DeleteRoute(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete route"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
