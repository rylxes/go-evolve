package utils

import (
	"github.com/gin-gonic/gin"
	"go-postgres/domain"
	"strconv"
)

func GeneratePaginationFromRequest(c *gin.Context) domain.Pagination {
	// Initializing default
	//	var mode string
	limit := 10
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return domain.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}
