package controllers

import (
	"net/http"

	UserRequests "github.com/saegus/test-technique-romain-chenard/api/dto/requests"
	Service "github.com/saegus/test-technique-romain-chenard/api/services"
	"github.com/saegus/test-technique-romain-chenard/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService Service.UserServiceInterface
}

func NewAuthCtrl() *AuthController {
	return &AuthController{
		userService: Service.NewUserSrv(),
	}
}

func (controller *AuthController) HandleSignup(c *gin.Context) {
	var newUserReceived UserRequests.SignupRequest

	if err := c.ShouldBind(&newUserReceived); err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err := utils.PasswordConstrainsValidator(newUserReceived.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	
	recordedUser, err := controller.userService.CreateUserSrv(newUserReceived)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": recordedUser})
}

func (controller *AuthController) HandleSignin(c *gin.Context){
	var signinInfo UserRequests.LoginRequest

	if err := c.ShouldBind(&signinInfo); err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := controller.userService.LoginSrv(signinInfo)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, userResponse)
}

func (controller *AuthController) Verify(c *gin.Context){
	id, _ := c.Get("user_id")
	email, _ := c.Get("user_email")
	c.JSON(http.StatusAccepted, gin.H{"id": id, "email": email})
}
