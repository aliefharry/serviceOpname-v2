package utils

import(
	"serviceOpname-v2/config/entity/helper"
	"strconv"
	// "fmt"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationFromRequest(c *gin.Context) helper.Pagination {
	// initializing default
	//var mode string

	limit := 10
	page := 0
	sort := "id asc"
	
	query := c.Request.URL.Query()
	// fmt.Println("\n \n query value ", fmt.Sprintf("%#v", query))

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