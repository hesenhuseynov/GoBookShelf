package handler

import (
	"GoBookShelf/dto"
	"GoBookShelf/pkg/models"
	"GoBookShelf/pkg/service"
	"GoBookShelf/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	jwtService  service.JWTService
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (c *UserHandler) Register(ctx *gin.Context) {
	var user dto.UserCreateRequest
	if err := ctx.ShouldBind(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"action": "Hashing password",
			"error":  err.Error(),
			"where":  "userHandler",
		}).Error("Failed to has password")
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_HASHPASSWORD_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	user.Password = string(hashedPassword)

	result, err := c.userService.RegisterUser(ctx.Request.Context(), user)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"action": "Registering user",
			"error":  err.Error(),
			"where":  "userHandler",
		}).Error("Failed to register user")

		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return

	}
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	ctx.JSON(http.StatusOK, res)

}

func (c *UserHandler) Login(ctx *gin.Context) {
	var req dto.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.userService.Verify(ctx.Request.Context(), req.Email, req.Password)
	if err != nil && !res {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, err := c.userService.GetUserByEmail(ctx.Request.Context(), req.Email)
	if err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := c.jwtService.GenerateToken(user.ID, user.Role)
	userResponse := models.Authorization{
		Token: token,
		Role:  user.Role,
	}

	response := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, userResponse)
	ctx.JSON(http.StatusOK, response)

	// res,err:=c.userService.V

}
