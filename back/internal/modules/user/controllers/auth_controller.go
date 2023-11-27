package controllers

import (
	"net/http"

	UserRequests "github.com/saegus/test-technique-romain-chenard/internal/modules/user/requests"
	UserService "github.com/saegus/test-technique-romain-chenard/internal/modules/user/services"
	"github.com/saegus/test-technique-romain-chenard/pkg/utils"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
	}
}

func (controller *Controller) HandleSignup(c *gin.Context) {
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

func (controller *Controller) HandleSignin(c *gin.Context){
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

func (controller *Controller) Verify(c *gin.Context){
	id, _ := c.Get("user_id")
	c.JSON(http.StatusAccepted, gin.H{"id": id})
}
