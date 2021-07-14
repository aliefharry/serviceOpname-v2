package controller

import(
	"net/http"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"serviceOpname-v2/service"
	"serviceOpname-v2/config/entity"
	"serviceOpname-v2/controller/dto"
	"serviceOpname-v2/config/entity/helper"
	"github.com/dgrijalva/jwt-go"

)

type OpnameController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Update(context *gin.Context)
}

type opnameController struct {
	opnameService service.OpnameService
	jwtService service.JWTService
}

func NewOpnameController(opnameServ service.OpnameService, jwtServ service.JWTService) OpnameController {
	return &opnameController{
		opnameService: opnameServ,
		jwtService: jwtServ,
	}
}

func(c *opnameController) getUserIDByToken(token string) string {
	Token, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["userID"])
	return id
}

func (c *opnameController) All(context *gin.Context) {
	var opnames []entity.Opname = c.opnameService.All()
	res := helper.BuildResponse(true, "OK", opnames)
	context.JSON(http.StatusOK, res)
}

// func (c *opnameController) All(context *gin.Context) {
// 	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
// 	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "3"))

// 	var opnames []entity.Opname = c.opnameService.All()

// 	paginator := pagination.Paging(&pagination.Param{
//         DB:      c,
//         Page:    page,
//         Limit:   limit,
//         OrderBy: []string{"id desc"},
//         ShowSQL: true,
//     }, &opnames)
// 	res := helper.BuildResponse(true, "OK", opnames)
// 	context.JSON(http.StatusOK, res)
// }

func (c *opnameController) FindById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var book entity.Opname = c.opnameService.FindById(id)
	if (book == entity.Opname{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", book)
		context.JSON(http.StatusOK, res)
	}
}

func (c *opnameController) Update(context *gin.Context) {
	var OpnameUpdDTO dto.OpnameUpdDTO
	errDTO := context.ShouldBind(&OpnameUpdDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["userID"])
	if c.opnameService.IsAllowedToEdit(userID, OpnameUpdDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			OpnameUpdDTO.UserID = id
		}
		result := c.opnameService.Update(OpnameUpdDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}