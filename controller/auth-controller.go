package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"serviceOpname-v2/service"
	"serviceOpname-v2/config/entity"
	"serviceOpname-v2/controller/dto"
	"serviceOpname-v2/config/entity/helper"
)

//authcontroller interface is contract what this controller can do
type AuthController interface {

	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
//isi service yg dibutuh
	authService service.AuthService
	jwtService service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService : authService,
		jwtService : jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context){
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context){
	var registerDTO dto.RegisDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failde to Process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email){
		response := helper.BuildErrorResponse("Failde to Process request", "Duplicate email", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}else{
		createdUser := c.authService.CreateUser(registerDTO)
		// token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		// createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}