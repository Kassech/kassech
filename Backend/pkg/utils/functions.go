package utils

import (
	"fmt"
	"math"
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

// Haversine calculates the great-circle distance between two points
// on the Earth (specified in decimal degrees) using the Haversine formula.
// Returns distance in kilometers.
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371 // Earth's mean radius in kilometers

	// Convert degrees to radians
	degToRad := func(deg float64) float64 { return deg * math.Pi / 180 }

	lat1Rad := degToRad(lat1)
	lon1Rad := degToRad(lon1)
	lat2Rad := degToRad(lat2)
	lon2Rad := degToRad(lon2)

	// Differences in coordinates
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	// Haversine formula components
	a := math.Pow(math.Sin(dLat/2), 2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Pow(math.Sin(dLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Calculate the distance
	distance := earthRadius * c

	return distance
}
