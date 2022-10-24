package routes

import (
	"net/http"

	"github.com/ddld93/promon/project/src/controller"
	"github.com/ddld93/promon/project/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type ProjectRoute struct {
	ProjectCtrl *controller.DB_Connect
}
type CustomResponse struct {
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
	Error   string      `json:"error"`
}



func (ctrl *ProjectRoute) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var project model.Project
		//validate the request body
		if err := c.BindJSON(&project); err != nil {
			c.JSON(http.StatusBadRequest, CustomResponse{
				Message: "Error parsing request body",
				Payload: nil,
				Error:   err.Error(),
			})
			return
		}
	
		res, err := ctrl.ProjectCtrl.AddOne(project)

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

func (ctrl *ProjectRoute) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := ctrl.ProjectCtrl.GetAll()
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

func (ctrl *ProjectRoute) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		Id := c.Param("id")
		objId, _ := primitive.ObjectIDFromHex(Id)
		res, err := ctrl.ProjectCtrl.GetOne(objId)
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

func (ctrl *ProjectRoute) UpdateOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		Id := c.Param("id")
		objId, _ := primitive.ObjectIDFromHex(Id)
		var project model.Project
		//validate the request body
		if err := c.BindJSON(&project); err != nil {
			c.JSON(http.StatusBadRequest, CustomResponse{
				Message: "Error parsing request body",
				Payload: nil,
				Error:   err.Error(),
			})
			return
		}
		res := ctrl.ProjectCtrl.Update(objId,project)
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

func (ctrl *ProjectRoute) DeleteOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		Id := c.Param("id")
		objId, _ := primitive.ObjectIDFromHex(Id)
		
		res := ctrl.ProjectCtrl.Delete(objId)
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