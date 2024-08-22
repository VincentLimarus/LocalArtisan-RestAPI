package services

import (
	"fmt"
	"localArtisans/controllers/helpers"
	"localArtisans/models/outputs"
	"localArtisans/models/requestsDTO"
	"localArtisans/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

func GetCarts(c *gin.Context){
	var GetAllCartsRequestDTO requestsDTO.GetAllCartsRequestDTO
	GetAllCartsRequestDTO.Page, GetAllCartsRequestDTO.Limit, GetAllCartsRequestDTO.OrderBy, GetAllCartsRequestDTO.OrderType = utils.PaginationHandler(GetAllCartsRequestDTO.Page, GetAllCartsRequestDTO.Limit, GetAllCartsRequestDTO.OrderBy, GetAllCartsRequestDTO.OrderType)
	if err := c.ShouldBindWith(&GetAllCartsRequestDTO, binding.Form); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllCarts(GetAllCartsRequestDTO)
	c.JSON(code, output)
}

func GetCartByUserID(c *gin.Context){
	cartID := c.Param("id")

	if _, err := uuid.Parse(cartID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllCartsByUserID(cartID)
	c.JSON(code, output)
}

func GetCartByID(c *gin.Context){
	cartID := c.Param("id")

	if _, err := uuid.Parse(cartID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetCartByID(cartID)
	c.JSON(code, output)
}

func CreateCart(c *gin.Context){
	var CreateCartRequestDTO requestsDTO.CreateCartRequestDTO
	if err := c.ShouldBindJSON(&CreateCartRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.CreateCart(CreateCartRequestDTO)
	c.JSON(code, output)
}

func DeleteCart(c *gin.Context){
	var DeleteCartRequestDTO requestsDTO.DeleteCartRequestDTO
	if err := c.ShouldBindJSON(&DeleteCartRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.DeleteCart(DeleteCartRequestDTO)
	c.JSON(code, output)
}

func AuthCartService(router *gin.RouterGroup) {
	router.GET("/carts", GetCarts)
	router.GET("/carts/user/:id", GetCartByUserID)
	router.GET("/cart/:id", GetCartByID)
	router.POST("/cart/create", CreateCart)
	router.POST("/cart/delete", DeleteCart)
}