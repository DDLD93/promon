package routes

import (
	"net/http"

	"github.com/ddld93/promon/contractor/src/controller"
	"github.com/ddld93/promon/contractor/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContractorRoute struct {
	ContractorCtrl *controller.DB_Connect
}
type CustomResponse struct {
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
	Error   string      `json:"error"`
}

func (ctrl *ContractorRoute) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var contractor model.Project
		//validate the request body
		if err := c.BindJSON(&contractor); err != nil {
			c.JSON(http.StatusBadRequest, CustomResponse{
				Message: "Error parsing request body",
				Payload: nil,
				Error:   err.Error(),
			})
			return
		}

		res, err := ctrl.ContractorCtrl.AddOne(contractor)

		if err != nil {
			c.JSON(http.StatusBadRequest, CustomResponse{
				Message: "An Error occured adding to database",
				Payload: nil,
				Error:   err.Error()})
			return
		}
		c.JSON(http.StatusCreated, CustomResponse{
			Message: "success",
			Payload: res,
			Error:   "",
		})
		return
	}
}

func (ctrl *ContractorRoute) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := ctrl.ContractorCtrl.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, CustomResponse{
				Message: "An error occured querying database",
				Payload: nil,
				Error:   err.Error()})
			return
		}
		c.JSON(http.StatusOK, CustomResponse{
			Message: "success",
			Payload: res,
			Error:   "",
		})

	}
}

func (ctrl *ContractorRoute) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		Id := c.Param("id")
		objId, _ := primitive.ObjectIDFromHex(Id)
		res, err := ctrl.ContractorCtrl.GetOne(objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, CustomResponse{
				Message: "An error occured querying database",
				Payload: nil,
				Error:   err.Error()})
			return
		}
		c.JSON(http.StatusCreated, CustomResponse{
			Message: "success",
			Payload: res,
			Error:   "",
		})

	}
}

func (ctrl *ContractorRoute) UpdateOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		Id := c.Param("id")
		objId, _ := primitive.ObjectIDFromHex(Id)
		var contractor model.Project
		//validate the request body
		if err := c.BindJSON(&contractor); err != nil {
			c.JSON(http.StatusBadRequest, CustomResponse{
				Message: "Error parsing request body",
				Payload: nil,
				Error:   err.Error(),
			})
			return
		}
		res := ctrl.ContractorCtrl.Update(objId, contractor)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, CustomResponse{
		// 		Message: "An error occured querying database",
		// 		Payload: nil,
		// 		Error:   err.Error()})
		// 	return
		// }
		c.JSON(http.StatusCreated, CustomResponse{
			Message: "success",
			Payload: res,
			Error:   "",
		})

	}
}

func (ctrl *ContractorRoute) DeleteOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		Id := c.Param("id")
		objId, _ := primitive.ObjectIDFromHex(Id)

		res := ctrl.ContractorCtrl.Delete(objId)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, CustomResponse{
		// 		Message: "An error occured querying database",
		// 		Payload: nil,
		// 		Error:   err.Error()})
		// 	return
		// }
		c.JSON(http.StatusCreated, CustomResponse{
			Message: "success",
			Payload: res,
			Error:   "",
		})

	}
}
