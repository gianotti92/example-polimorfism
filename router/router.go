package router

import (
	"awesomeProject/interact"
	"awesomeProject/router/controller"
	"github.com/gin-gonic/gin"
)

var engine = gin.Default()

func StartDefaultEngine() error {

	buyerClient := controller.ClientController{
		ClientType: "buyer",
		Interact: interact.BuyerInteract{},
	}

	sellerClient := controller.ClientController{
		ClientType: "seller",
		Interact: interact.SellerInteract{},
	}


	engine.GET("/ping", controller.Ping)
	engine.POST("/buyer", buyerClient.Client)
	engine.POST("/seller", sellerClient.Client)

	err := engine.Run()

	return err
}


