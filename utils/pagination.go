package utils

import(
	"serviceOpname-v2/config/entity/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationFromRequest(c *gin.Context) helper.Pagination {
	// initializing default
	//var mode string

	limit := 10
	page := 1
	sort := "id desc"
	
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

	return helper.Pagination{
		Limit: limit,
		Page: page,
		Sort: sort,
	}
}