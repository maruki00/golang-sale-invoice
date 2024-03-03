package controllers

import (
	interactors "saleinvoice/App/Application/Interactors"
	"saleinvoice/App/Application/helpers"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	interactor interactors.CategoryInteractor
}

func NewCategoryController(interactor interactors.CategoryInteractor) *CategoryController {
	return &CategoryController{
		interactor: interactor,
	}
}

func (c *CategoryController) Index(ctx *gin.Context) {
	category := c.interactor.Index()
	ctx.JSON(200, category)
}

func (controller *CategoryController) Store(ctx *gin.Context) {
	resp := controller.interactor.Store(ctx.Request)
	helpers.Success(ctx, resp, 200)
	// helpers.category := factories.CategoryFactory(ctx.Request.Body)
}
