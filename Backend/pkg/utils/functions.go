package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
)

func GetForceDeleteFromHeader(c *gin.Context) (bool, error) {
	roles, ok := c.Get("role")
	if !ok {
		return false, nil
	}
	isAdmin := utils.Contains(roles.([]string), "Admin")
	if isAdmin {
		forceDeleteStr := c.Query("force")
		fmt.Println("forceDeleteStr:", forceDeleteStr)
		if forceDeleteStr != "" {
			forceDelete, err := strconv.ParseBool(forceDeleteStr)
			if err != nil {
				return false, nil
			}
			return forceDelete, nil
		}
	}
	return false, nil
}

func GetPageFromQuery(c *gin.Context) (int, error) {
	pageStr := c.Query("page")
	if pageStr == "" {
		return 1, nil
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, err
	}
	if page < 1 {
		return 0, fmt.Errorf("invalid page number: %d", page)
	}
	return page, nil
}

func GetPerPageFromQuery(c *gin.Context) (int, error) {
	perPageStr := c.Query("per_page")
	if perPageStr == "" {
		return 10, nil
	}
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		return 0, err
	}
	if perPage < 1 {
		return 0, fmt.Errorf("invalid per page number: %d", perPage)
	}
	return perPage, nil
}
