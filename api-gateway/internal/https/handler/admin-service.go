package handler 

import (
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"

	"github.com/gin-gonic/gin"
)

// @Router /admin/departments [post]
// @Summary CREATE DEPARTMENT	
// @Security  		BearerAuth
// @Description This method creates department
// @Tags AUTH
// @Accept json
// @Produce json
// @Param department body models.RegisterUserRequest true "User"
// @Success 200 {object} models.RegisterUserResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) RegisterUser(c *gin.Context) {

	req := pb.RegisterUserRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("RegisterUser: JSON binding error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.RegisterUser(ctx, &req)
	if err != nil {
		logger.Error("RegisterUser: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("RegisterUser: Successfully registered user - ", resp.UserName)
	c.JSON(200, resp)
}
