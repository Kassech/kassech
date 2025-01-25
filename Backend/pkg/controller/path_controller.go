package controller

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PathController struct {
	Service *service.PathService
}

// CreatePath handles HTTP requests to create a new path
func (pc *PathController) CreatePath(c *gin.Context) {
	var path models.Path
	if err := c.ShouldBindJSON(&path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPath, err := pc.Service.CreatePath(&path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdPath)
}

// DeletePath handles HTTP requests to delete a path
func (pc *PathController) DeletePath(c *gin.Context) {
	pathID := c.Param("id")
	pathIDUint, err := utils.StringToUint(pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	deletedPath, err := pc.Service.DeletePathByID(pathIDUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedPath)
}

// FindPathByID handles HTTP requests to get a path by ID
func (pc *PathController) FindPathByID(c *gin.Context) {
	pathID := c.Param("id")
	pathIDUint, err := utils.StringToUint(pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	path, err := pc.Service.FindPathByID(pathIDUint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, path)
}

// FindPathByID handles HTTP requests to get a path by ID
func (pc *PathController) FindPathsByStationID(c *gin.Context) {
	pathID := c.Param("id")
	pathIDUint, err := utils.StringToUint(pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Station ID"})
		return
	}
	path, err := pc.Service.FindPathsByStationID(pathIDUint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, path)
}

// GetAllPaths handles HTTP requests to get all paths
func (pc *PathController) GetAllPaths(c *gin.Context) {
	page, err := utils.GetPageFromQuery(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	perPage, err := utils.GetPerPageFromQuery(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	search := c.Query("search")
	paths, total, err := pc.Service.GetAllPaths(page, perPage, search)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  paths,
		"total": total,
	})
}
